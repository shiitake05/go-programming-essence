// 可変個引数による関数の実装（型が固定となる）。
// package main

// import "fmt"

// func doSomething(args ...string) {
// 	for _, arg := range args {
// 		fmt.Println(arg)
// 	}
// }

// func main() {
// 	doSomething("Hello", "World!", "Golang")
// }

// Functional Options Patternを使用することで開発者がその関数の利用者に対して、開発者が限定した型のみを可変個で渡せる関数を提供できる

// 接続タイムアウトやロガーを追加
package main

import (
	"FunctionalOptionsPattern/server"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.New("localhost", 8888, server.WithTimeout(time.Minute), server.WithLogger(logger))
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
