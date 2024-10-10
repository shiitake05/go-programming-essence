// http://localhost:8080/ でHello, World!
package main

import (
	"io"
	"net/http"
	"os"
)

// 関数のシグネクチャが同じであれば引数として渡すことも可能
// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, World!")
// }
// func main() {
// http.Handle("/", myHandler)
// http.ListenAndServe(":8080", nil)
// }

// Goのメソッド値を使うことでhttp.HandleFuncにメソッドを渡すことができる
// これによってデータベース接続をdbに設定し、handleの中でそのdbを使った処理を実装することが可能
// type myContext struct{
// db *sql.DB
// }
// func (m *myContext) handle(w http.ResponseWriter, r *http.Request) {
// dbを使った処理
// }
// func main() {
// myctx := NewMyContext()
// http.HandleFunc("/", myctx.handle)
// ...
// }

func main() {
	// http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {	// 下と等価
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 	fmt.Fprint(w, "Hello, World!")

		// GETとPOSTを区別したい場合はr.Methodを確認する
		// 	switch r.Method {
		// 	case http.MethodGet:
		// 		// GET時の処理
		// 		fmt.Fprint(w, "GET")
		// 	default:
		// 	}

		// http.ResponseWriterはio.Writerであるため、Goのio.Writerを受け取る多くの関数を呼び出すことができる
		// contents := forecast()
		// w.Write([]byte(contents))

		// io.Copyでファイルの内容を出力
		f, err := os.Open("/path/to/content.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}

// func forecast() string {
// 	return "晴れ"
// }
