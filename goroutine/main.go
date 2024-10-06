package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(msg string) {
	println(msg)
}

type Test struct {
	name string
}

// goroutineはGoのランタイムで管理された軽量スレッド
// 重たい処理を並行して行える
func main() {
	// message := "hi"
	// go sendMessage(message) // 関数呼び出しの前に「go」を付けると、その関数をゴルーチンとして実行する
	// message = "bye"         // 呼び出される変数は変わらない

	// ただし、無名関数の場合は書き換えの方が先に実行される可能性がある
	message := "hi"
	go func() {
		sendMessage(message)
	}()
	message = "bye"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
	// go run -race main.goでrace condition(goroutineとその呼び出し元との間でデータ競合が起きていること)を検出

	// goroutineが実行中にも関わらず、main関数の終了とともに強制終了されてしまうこともある
	// そのため、goroutineの終了を待つためにsyncパッケージを使う
	// var wg sync.WaitGroup
	// wg.Add(1) // リファレンスカウンタを+1する
	// go func() {
	// 	defer wg.Done() // リファレンスカウンタを-1する
	// 	// 重い処理
	// }()
	// // 重い処理
	// wg.Wait() // リファレンスカウンタが0になるまで待つ

	// 以下のコードを実行するとランダムな数字が羅列される
	// これは、goroutineが実行されるまでにforループが終了してしまうため
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}()
	// }
	// wg.Wait()

	// 以下のように書き換えることで、goroutineが正しいループ変数を参照するようになる
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()

	// 別方法
	tests := []Test{
		{name: "test1"},
		{name: "test2"},
		{name: "test3"},
	}
	for _, tt := range tests {
		wg.Add(1)
		go func(tt *Test) {
			defer wg.Done()
			fmt.Println(tt.name)
		}(&tt)
	}
	wg.Wait()

	// 別方法
	for i := range tests {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(tests[i].name)
		}()
	}
	wg.Wait()

	// データ保護をするためにはsync.Mutexを使って保護する必要がある
	// 使用していない場合(2000が表示されないことがある)
	n := 0
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			n++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			n++
		}
	}()
	wg.Wait()
	fmt.Println(n)

	// 使用している場合(2000が表示される)
	n2 := 0
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n2++
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n2++
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(n)
}
