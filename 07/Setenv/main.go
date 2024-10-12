package main

import (
	"fmt"
	"os"
)

// doSomethingは環境変数に依存した処理を行う関数です。
func doSomething2() error {
	// 環境変数DATABASE_URLを取得
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URLが設定されていません")
	}
	// ここでデータベース接続などの処理を行うことができます
	fmt.Println("接続先データベース:", dbURL)
	return nil
}

func main() {
	// 環境変数DATABASE_URLを設定
	os.Setenv("DATABASE_URL", "localhost")
	// doSomething関数を呼び出す
	if err := doSomething2(); err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	fmt.Println("処理が成功しました。")
}
