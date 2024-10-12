package main

import (
	"fmt"
)

func doSomething(data string) string {
	// ここに処理を追加
	return fmt.Sprintf("Processed: %s", data)
}

func main() {
	// テスト用のデータを用意
	testData := "example input"
	result := doSomething(testData)
	fmt.Println(result)
}
