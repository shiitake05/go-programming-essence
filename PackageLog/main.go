// logパッケージは簡単なログの仕組みを提供する
package main

import (
	"log"
	"os"
)

func main() {
	filename := "app.log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)           // ログの出力先をファイルに変更
	log.Println("app started") // ログを出力
}
