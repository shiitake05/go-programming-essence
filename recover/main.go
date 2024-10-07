package main

import (
	"fmt"
)

// 配列によって境界外アクセスによってランタイムでpanicが発生するがrecoverで復帰する
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	var a [2]int
	n := 2
	println(a[n])
}
