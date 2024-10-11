// generatorに対して停止指示を出せるようにしている
// 停止指示を行った後、後処理が必要な場合は、さらにquitから応答をもらえるようにしている
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	ch := generator("Hi!", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "Bye!"
	fmt.Printf("Generator says %s", <-quit)
}

func generator(msg string, quit <-chan string) <-chan string {
	ch := make(chan string)
	go func() { // 匿名関数goroutine
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// 何もしない
			case <-quit:
				fmt.Println("Goroutine done")
				return
			}
		}
	}()
	return ch
}
