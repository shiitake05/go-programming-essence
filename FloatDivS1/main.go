// 引数を割り算し浮動小数点の値として結果を表示するプログラム
// (FloatDivより短く表記できているが、発生し得るエラーをユーザーに伝えられない)
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintln(os.Stderr, e)
		}
	}()
	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)
	fmt.Println(a / b)
}
