// goに標準で用意されているフラグパーサ
// フラグを扱うための基本的な機能のみを提供し、コマンドライン引数として、フラグと入力文字列だけを扱うようなプログラムを作る際に便利
// これは「something 255」が表示される
package main

import (
	"flag"
	"fmt"
)

func main() {
	// max := flag.Int("max", 255, "max value")
	// name := flag.String("name", "something", "my name")
	// flag.Parse()	// この実行で引数が解析され、不正な引数が渡された場合にプログラムを中断して使用方法を表示する

	// println(*name, *max)

	// flag.Intはint型のポインタを返すが、ポインタでない値が欲しい時はIntVarを使用する
	// IntとIntVar, StringとStringVar, BoolとBoolVarがある
	var name string
	var max int

	flag.IntVar(&max, "max", 255, "max value")
	flag.StringVar(&name, "name", "", "my name")
	flag.Parse()

	// flagパッケージがフラグを解析したあとの残りの引数は、flag.Argsから得ることができる
	for _, arg := range flag.Args() {
		fmt.Println(arg)
	}
}
