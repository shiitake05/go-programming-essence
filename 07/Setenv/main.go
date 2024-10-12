package main

import (
	"fmt"
	"os"
)

func doSomething2() error {
	// 環境変数DATABASE_URLを取得
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URLが設定されていません")
	}
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
