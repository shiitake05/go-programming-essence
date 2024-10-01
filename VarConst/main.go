package main

import "fmt"

func main() {
	var n int = 1
	fmt.Println(n)
	// n = "foo" // コンパイルエラー
	n = 1

	// var n string
	//n = "foo" // コンパイルエラー

	x := 1   // var x int = 1 と同じ
	y := 1.2 // varを省略可能
	n1 := 1 + (x+2)*int(y)
	n2 := 1.2 + (float64(x)+2)*float64(y) // nはfloat64型の変数
	fmt.Println(n1, n2)

	s := "hello, " + "world" // 文字列は+で結合できる
	fmt.Println(s)

	const x2 = 1 // 定数の宣言
	y2 := 1
	f := 1.2 + (x2+2)*float64(y2) // 定数は使われる箇所で型が決まる
	fmt.Println(f)

	const n3 = 1
	x3 := 1 + n3
	y3 := 1.2 + n3
	fmt.Println(x3, y3) // コンパイルエラーにならない

}
