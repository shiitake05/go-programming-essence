// Goはエラーを都度処理する文法であるため、ネットワークの処理を書くのに向いている
// このサーバアプリケーションは接続数が制限されていないがgoroutineは軽量スレッドとして動作するためパフォーマンスが劣化しない
package main

import (
	"log"
	"net"
)

func handleData(conn net.Conn) {
	defer conn.Close()

	var b [512]byte
	for {
		n, err := conn.Read(b[:])
		if err != nil {
			break
		}
		log.Println(string(b[:n]))
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleData(conn)
	}
}
