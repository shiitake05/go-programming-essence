package main

import "fmt"

func main() {
	s := "Hello"
	name := " user!"

	s += name

	fmt.Println(s)
	fmt.Printf("%c\n", s[0]) // Hを表示

	// バイト列の変換
	b := []byte(s)
	b[0] = 'h'
	fmt.Println(b)

	s = string(b)
	fmt.Println(s) // hello user!

	s1 := "こんにちわ世界"
	rs := []rune(s1)
	rs[4] = 'は'     // こんにちわ世界をこんにちは世界に変更
	s1 = string(rs) // イミュータブルなので再代入が必要
	fmt.Println(s1) // こんにちは世界

	var s2 [4]byte
	s3 := s2[:] // s3はスライスとなる
	fmt.Println(s3)

	// 複数行の文字列
	var content = `複数行の
	文章からなる
	テキストです。
	`

	var content2 = `
	複数行の
	文章からなる
	テキストです。
	`

	fmt.Println(content)
	fmt.Println(content2)
}
