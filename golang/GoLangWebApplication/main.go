package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {

	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	// ゴルーチンが異常終了したときに ctx をキャンセルできるようにする
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// ゴルーチンの終了と error を受け取るためのチャンネル。
	// バッファを 1 にしないと、受信前に run が return した場合に送信側がリークする
	errCh := make(chan error, 1)

	// 別のゴルーチンでサーバーを起動
	go func() {
		if err := s.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to close: %+v", err)
			errCh <- err
			// 起動に失敗した場合は自分で ctx をキャンセルしないと
			// 下の <-ctx.Done() から抜けられない
			cancel()
			return
		}
		errCh <- nil
	}()

	// チャンネルからの通知（）終了通知を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
		return err
	}
	// ゴルーチンの終了を待ち、その error を返す
	return <-errCh
}
