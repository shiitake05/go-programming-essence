package main

import (
	"net/http"
	"path"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".static")))                                     // ファイルをサーブしたい場合に使う
	http.Handle("/public/", http.FileServer(http.Dir("./static")))                             // これは/public/index.htmlというリクエストに対して、ファイルサーバが./public/static/index.htmlというファイルを探してしまう
	http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("./static")))) // SriptPrefixの第一引数は第二引数に対して作用するのではなく、リクエストURIに対して作用する

	// FileServerを使う場合、Content-Typeが自動で設定される
	//  GOが用意するMIMEタイプに該当しないファイルをセーブしたい場合はファイルのサーブ処理を書く必要がある
	fileserver := http.StripPrefix("/public", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) == ".xls" {
			w.Header().Set("Content-Type", "application/vnd.ms-excel")
		}
		fileserver.ServeHTTP(w, r)
	})

	// カスタマイズしたい拡張子がたくさんある場合はmapを用意する
	mimemap := map[string]string{
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"ppt":   "application/vnd.ms-powerpoint",
		"pptx":  "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ext := path.Ext(r.URL.Path)
		if mime, ok := mimemap[ext]; ok {
			w.Header().Set("Content-Type", mime)
		}
		fileserver.ServeHTTP(w, r)
	})
}
