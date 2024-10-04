package main

import "fmt"

// tyoeで型に名前をつけることができる
type MyString string

func main() {
	// MyStringはstringのように振る舞う
	var m MyString
	m = "foo"
	fmt.Println(m)

	// 元のstring型に戻す
	s := string(m)
	fmt.Println(s)
}
