BINARY := learninggo
BIN_DIR := bin

.DEFAULT_GOAL := help
.PHONY: run build test fmt vet tidy check clean help

## run: プログラムを実行
run:
	go run .

## build: バイナリを bin/ に出力
build:
	go build -o $(BIN_DIR)/$(BINARY) .

## test: テストを実行
test:
	go test ./...

## fmt: コードを整形
fmt:
	go fmt ./...

## vet: 静的解析
vet:
	go vet ./...

## tidy: 依存関係を整理
tidy:
	go mod tidy

## check: fmt → vet → test を順に実行
check: fmt vet test

## clean: ビルド成果物を削除
clean:
	rm -rf $(BIN_DIR)

## help: ターゲット一覧を表示
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/^## /  /'
