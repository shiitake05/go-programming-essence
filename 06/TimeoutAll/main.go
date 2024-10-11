// 全体をタイムアウト処理
// goroutineを停止する処理はないが、タイムアウトによって5秒後に終了する
// channelの送信がブロッキングするため、ループが回り続けることはないが、gotoutineは生き続ける
// よって、generatorを必要しなくなった後は、goroutineが生き続けるため、goroutineリークが発生する
// timeout処理の二つのプログラムはgoroutineリークが発生することはないが、ライブラリとして何度か使う場合はgoroutineリークが発生する
// 安全に使用したい場合は、StopInで実装したようにする
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hi!")
	timeout := time.After(5 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("5s Timeout!")
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
