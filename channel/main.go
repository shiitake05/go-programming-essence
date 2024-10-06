package main

import (
	"fmt"
)

func server(ch chan string) {
	defer close(ch) // channelを閉じる
	ch <- "one"     // channelにデータを送信
	ch <- "two"
	ch <- "three"
}

// channelはgoroutine間でデータをやり取りするための機能
func main() {
	var s string

	ch := make(chan string) // channelを作成
	go server(ch)

	s = <-ch
	fmt.Println(s) // one
	s = <-ch
	fmt.Println(s) // two
	s = <-ch
	fmt.Println(s) // three
}
