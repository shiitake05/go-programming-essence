package main

import (
	"fmt"
)

// switch文のようにchannelを読み込むことができる
// 効率が良い
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case v1 := <-ch1:
		fmt.Println("ch1から受信: ", v1)
	case v2 := <-ch2:
		fmt.Println("ch2へ送信: ", v2)
	default:
		fmt.Println("default")
	}

	// 重たい処理やブロッキングする処理を実行しつつ、main関数側でそれらのgoroutineからデータ入力がない間にも他のことをしたいケースで便利
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("ch1から受信: ", v1)
		case v2 := <-ch2:
			fmt.Println("ch2から受信: ", v2)
		default:
			fmt.Println("default")
			return
		}
	}
}
