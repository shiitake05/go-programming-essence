// goroutineは軽量であるため、使用し過ぎてしまうが、生成しすぎるとgoroutineの中で行われるI/Oがボトルネックになってしまう
// このときに必要なのが、用途に合わせて頻度や速度を調整する「スロットリング」で、サーバが処理できる程度に調整/制御を行う
// 以下の例は600秒かかってしまう処理
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func downloadJSON(u string) {
// 	// download JSON
// 	println(u)
// 	time.Sleep(time.Second * 6)
// }

// func main() {
// 	before := time.Now()
// 	for i := 0; i <= 100; i++ {
// 		u := fmt.Sprintf("https://example.com/api/users?id=%d", i)
// 		downloadJSON(u)
// 	}

// 	fmt.Printf("%v\n", time.Since(before))
// }

// これをgotoutineを生成して並列にする
// 100個のgoroutineが一気に生成されるため、6秒近くで終了する
// しかし、スロットリングしていないため、サーバに100個のリクエストが送信されてしまう
// そのため、goroutineの生成を20個に制限し、バッファ付きのchanが、バッファを超える要素を送信しようとした時にブロックする特性を利用する
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func downloadJSON(u string) {
// 	// download JSON
// 	println(u)
// 	time.Sleep(time.Second * 6)
// }

// func main() {

// 	var wg sync.WaitGroup
// 	before := time.Now()
// 	for i := 0; i <= 100; i++ {
// 		wg.Add(1)
// 		i := i
// 		go func() {
// 			defer wg.Done()

// 			u := fmt.Sprintf("https://example.com/api/users?id=%d", i)
// 			downloadJSON(u)
// 		}()
// 	}
// 	wg.Wait()

// 	fmt.Printf("%v\n", time.Since(before))
// }

// 計算上30秒になる
package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	// download JSON
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()

	limit := make(chan struct{}, 20)
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)

		i := i
		go func() {
			// limit が20よりも多いとブロックする
			limit <- struct{}{}
			defer wg.Done()
			u := fmt.Sprintf("https://example.com/api/users?id=%d", i)
			downloadJSON(u)
			<-limit // limit から抜き出す
		}()
	}
	wg.Wait()

	fmt.Printf("%v\n", time.Since(before))
}
