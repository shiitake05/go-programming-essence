package main

import (
	"fmt"
	"reflect"
)

func main() {
	// interface{}型は任意の型を代入できる
	var v interface{}
	v = 1
	fmt.Println(v)
	v = "こんにちは世界"
	fmt.Println(v)

	// 型アサーションによって型を取得することができる
	v = 1
	n := v.(int)

	v = "こんにちは世界"
	s := v.(string)

	fmt.Println(n, s)

	// 間違った型アサーションを行うとpanicが発生する
	// 正しい型アサーションができなかったことを確認するには以下のように実行する
	// s, ok := v.(string)
	// if !ok {
	// 	fmt.Println("vはstringではない")
	// } else {
	// 	fmt.Printf("vは文字列 %q です\n", s)
	// }

	// 1行で書くと以下の通り
	if s, ok := v.(string); !ok {
		fmt.Println("vはstringではない")
	} else {
		fmt.Printf("vは文字列 %q です\n", s)
	}

	type V int
	var v2 V = 123
	PrintDetail(v2) // 知らない型が表示される（type定義された型でも扱えるようにするためにはreflectパッケージを使う必要がある）
}

// 型スイッチによってどんな値でも引数として受け取れ、処理できる関数を実装する
// func PrintDetail(v interface{}) {
// 	switch t := v.(type) {
// 	case int, int32, int64:
// 		fmt.Println("int/int32/int64 型:", t)
// 	case string:
// 		fmt.Println("string 型:", t)
// 	default:
// 		fmt.Println("知らない型")
// 	}
// }

func PrintDetail(v interface{}) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/int32/int64 型:", v)
	case reflect.String:
		fmt.Println("string 型:", v)
	default:
		fmt.Println("知らない型")
	}
}
