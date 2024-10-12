package main

import (
	"fmt"
	"os"
)

// doSomethingは文字列を受け取り、そのまま返す関数です。
func doSomething(s string) string {
	return s
}

func main() {
	// コマンドライン引数からファイル名を取得
	if len(os.Args) < 2 {
		fmt.Println("ファイル名を指定してください。")
		return
	}

	fileName := os.Args[1]

	// ファイルを読み込む
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("ファイルの読み込みに失敗しました: %v\n", err)
		return
	}

	// doSomething関数を呼び出し、結果を表示
	result := doSomething(string(data))
	fmt.Println("結果:", result)
}
