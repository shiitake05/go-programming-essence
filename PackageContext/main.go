// 非同期アプリケーションを開発するときに重要となるパッケージ
// goroutineを起動した側がgoroutineを中断したいときに使える
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("goroutine: 処理")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// ctx, cancel := context.WithCancel(context.Background())
	// go f(ctx, &wg)

	// time.Sleep(3 * time.Second)
	// cancel()

	// タイムアウトの場合
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	go f(ctx, &wg)

	wg.Wait()
}
