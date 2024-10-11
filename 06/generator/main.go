// 非同期の並行処理
package main

import (
	"fmt"
	"time"
)

func generator(msg string) <-chan string { // chanによって受信専用であることを明示している
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}

func main() {
	ch := generator("Hello!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
