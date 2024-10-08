package main

import (
	"fmt"
)

type F struct {
	Name string
	Age  int
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
	}

	fmt.Printf("%v\n", f) // &{John 20}

	// フィールド名とその内容が出力
	fmt.Printf("%+v\n", f) // &{Name:John Age:20}

	// struct名を出力
	fmt.Printf("%#v\n", f) // &main.F{Name:"John", Age:20}

	// 型名を出力したい場合は%Tを使用する
	fmt.Printf("%T\n", f)  // *main.F
	fmt.Printf("%T\n", *f) // main.F
}
