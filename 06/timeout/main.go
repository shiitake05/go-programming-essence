// generatorのループ内の処理が1秒以上かかると強制的に終了する
// ループ1つのタイムアウトではなく、全体のタイムアウトさせたい場合はTimeoutAllで実装している
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hi!")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Waited too lang!")
			return
		}
	}
}

func generator(msg string) <-chan string { // 受信専用のchannelを返す
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
