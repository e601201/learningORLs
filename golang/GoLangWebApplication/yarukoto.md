ステップ 1：とにかくサーバーを立てる

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello"))
    })
    log.Fatal(http.ListenAndServe(":8080", mux))
}

- 確認：curl localhost:8080/hello
- 学べること：ServeMux、ハンドラ関数の形 (w, r)

ステップ 2：User 型を作り、JSON で返す

- User 構造体と JSON タグ（json:"id"）を書く
- 固定の User を1つ作り、json.NewEncoder(w).Encode(...) で返す
- 確認：JSON が返ってくるか
- 学べること：構造体、タグ、Content-Type ヘッダー

ステップ 3：Store を作る（Mutex はまだ付けない）

- users map[int]User と nextID だけ持たせる
- NewStore() と Create()、List() を書く
- 試してほしいこと：make を書き忘れて、panic すること。

ステップ 4：Create と List のハンドラ

- POST /users：リクエストボディを json.NewDecoder(r.Body).Decode(&u) で受け取る
- GET /users：一覧を返す
- ここで初めて「作って、一覧で見る」がつながります。
- 確認：POST を数回 → GET で増えているか

ステップ 5：Get、Update、Delete を足す

1. GET /users/{id}：r.PathValue("id") と strconv.Atoi で ID を取る
2. DELETE /users/{id}：存在しなければ 404
3. PUT /users/{id}：最初は Name string で書いてみる

Update は最初に素直に書いてみるのがポイントです。 {"active":true} だけ送ると name が空になってしまいます。これを体験してから *string /
*bool（ポインタ）に直すと、「送られなかった」と「空文字が送られた」を区別するためにポインタを使う理由がよく分かります。

ステップ 6：重複をヘルパー関数にまとめる

ここまで書くと、同じコードが何度も出てきているはずです。

- JSON を返す処理 → writeJSON
- エラーを返す処理 → writeError
- ID を取り出す処理 → parseID

先に作るのではなく、3回くらい同じことを書いてから切り出すと、なぜ関数にまとめたいのかが身につきます。

ステップ 7：入力チェックを固める

- name が空なら 400
- DisallowUnknownFields() で未知のフィールドを弾く
- MaxBytesReader でボディのサイズを制限
- 確認：前回の「エラーになるケース」の表を curl で全部試す

ステップ 8：Mutex を付ける

- Store に sync.RWMutex を足し、各メソッドで Lock / RLock する
- 試してほしいこと：付ける前に次のようにすると、競合が起きることを確認できます。
go run -race .
# 別ターミナルで POST を並列に大量に送る
for i in $(seq 50); do curl -s -X POST localhost:8080/users -d '{"name":"x"}' & done
  -race を付けると、WARNING: DATA RACE が出て、Mutex が必要な理由が分かります。

ステップ 9：テストを書く

- httptest.NewRecorder() を使うと、サーバーを起動せずにハンドラ単体をテストできます。
- 正常系（CRUD の一連の流れ）→ 異常系（テーブル駆動テスト）の順で書く
- これまで curl で手動確認していたことを、自動化する

進めるときのコツ

- こまめに go vet と go fmt を実行する。 保存時に自動 Go 拡張機能を入れておくと楽です。
- 詰まったら見本と比べる。 ただし、写す前に一度エラーメッセージを読んでみてください。Go のエラーは比較的親切です。
- 完成したら、前に話した「internal/user に分ける」などのパッケージ分割に進むと、次の段階の学習になります。

どこかのステップで詰まったら、そのときのコードを見せてください。