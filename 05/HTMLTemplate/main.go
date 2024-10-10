// Goに標準で用意されたテンプレートエンジン
package main

import (
	"html/template"
	"log"
	"os"
)

// type User struct {
// 	Name string
// }

type User struct {
	Age  int
	Name string
}

type Employee struct {
	Name string
}

type Company struct {
	Employees []Employee
}

func main() {
	// Executeの基本形
	// tmpl := `{{.}}` // .はExecuteに渡された値、またはwith、rangeで束縛される値を示す
	// t := template.Must(template.New("").Parse(tmpl))
	// err := t.Execute(os.Stdout, "Hello, World!") // 文字列を渡す
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// struct
	// tmpl := `{{.Name}}` // .はExecuteに渡された値、またはwith、rangeで束縛される値を示す
	// t := template.Must(template.New("").Parse(tmpl))
	// user := User{Name: "Alice"}
	// err := t.Execute(os.Stdout, user) // User型の値を渡す
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// range/end(配列やスライスをExecuteに渡すと、rangeで繰り返し処理ができる)
	// tmpl := `{{range .}}` // .はExecuteに渡された値、またはwith、rangeで束縛される値を示す
	// tmpl += `
	// <p>{{.}}</p>{{end}}`
	// tmpl += `
	// {{index . 1}}`
	// t := template.Must(template.New("").Parse(tmpl))
	// values := []string{"Hello", "World"}
	// err := t.Execute(os.Stdout, values)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if/else/end(制御構文)
	// tmpl := `{{if gt .Age 20}}
	// {{.Name}} is older than 20
	// {{else}}
	// {{.Name}} is not older than 20
	// {{end}}`
	// t := template.Must(template.New("").Parse(tmpl))
	// // user := User{Age: 20, Name: "Alice"}
	// user := User{Age: 21, Name: "Bob"}
	// err := t.Execute(os.Stdout, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// with/else/end(ブロック内のスコープを変更)
	// tmpl := `{{with index .Employees 0}}
	// {{.Name}}
	// {{end}}`

	// withの値が空出会った場合にはelseを受け取れる
	// tmpl := `{{with .}}
	// {{.}}
	// {{else}}
	// Not Found
	// {{end}}`

	// withやrangeはコンテキスト内だけで有効な変数に値を代入
	// tmpl := `{{with $v := index .Employees 0}}
	// {{$v.Name}}
	// {{end}}`
	// t := template.Must(template.New("").Parse(tmpl))
	// company := Company{
	// 	Employees: []Employee{
	// 		{Name: "Alice"},
	// 		{Name: "Bob"},
	// 	},
	// }
	// err := t.Execute(os.Stdout, company)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// template.HTML
	// HTMLタグをそのまま出力する場合
	// tmpl := `<div>{{.}}</div>`
	// t := template.Must(template.New("").Parse(tmpl))
	// err := t.Execute(os.Stdout, template.HTML(`<b>HTML</b>`))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// template.JS
	// JavaScriptを出力する場合
	// tmpl := `<script>{{.}}</script>`
	// t := template.Must(template.New("").Parse(tmpl))
	// err := t.Execute(os.Stdout, template.JS(`alert("<script>1</script>")`))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// FuncMap
	// 自作の関数を追加
	// 日本語の日付をテンプレートで出力
	// t := template.New("").
	// 	Funcs(template.FuncMap{
	// 		"FormatDateTime": func(format string, d time.Time) string {
	// 			if d.IsZero() {
	// 				return ""
	// 			}
	// 			return d.Format(format)
	// 		}})
	// tmpl := `{{FormatDateTime "2006-01-02 15:04:05" .}}` // 実行するテンプレート
	// t = template.Must(t.Parse(tmpl))
	// err := t.Execute(os.Stdout, time.Now())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// template.Parse
	// 文字列を受け取ったり、ファイルを指定したりできる
	// サンプルコード
	// t := template.Must(template.New("").ParseFiles(
	// 	"template/index.html.tmpl",
	// 	"template/description.html.tmpl",
	// 	"template/login.html.tmpl",
	// ))
	// t := template.Must(template.New("").ParseGlob("template/*.tmpl"))
	// err := t.Execute(os.Stdout, "Hello, World!")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// define/template
	// テンプレートには名前をつけることができ、他のテンプレートから呼び出すことができる
	t := template.Must(template.New("").ParseGlob("template/*.tmpl"))
	err := t.ExecuteTemplate(os.Stdout, "index", "これは本文です") // index.html.tmplの各テンプレートの命令の先頭または末尾に-をつけることで改行を無効にできる
	if err != nil {
		log.Fatal(err)
	}
}
