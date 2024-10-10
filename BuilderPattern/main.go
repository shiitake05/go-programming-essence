// Functional Options Patternの対比される実装パターン
// Builder Patternはメソッドチェーンを用いることでBuildメソッドに与える初期値を決定する方法
package main

import (
	"BuilderPattern/server"
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
	svr := server.NewBuilder("localhost", 8888).Timeout(time.Minute).Logger(logger).Build()
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
