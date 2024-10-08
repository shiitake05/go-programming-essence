package main

import (
	"fmt"
	"os"
)

func main() {
	// fmtパッケージで数値を文字列に変換
	x := 1
	s1 := fmt.Sprintf("%d", x)
	s2 := fmt.Sprintf("%05d\n", x)
	fmt.Println(s1)
	fmt.Println(s2)

	// fmtパッケージでファイルに書き込む
	count := 5
	f, err := os.Create("output.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "COUNT %05d\n", count)
}
