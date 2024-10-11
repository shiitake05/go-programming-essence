// 「go build -o goroutine-test main.go」の後、「./goroutine-test」で実行
// 「ps a -o pid,nlwp,args | grep goroutine-test | grep -v grep」で確認
package main

import (
	"fmt"
	"sync"
)

// func doSomething(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(100 * time.Second)
// }

// 負荷が高くなるようにする
func doSomething(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	// time.Sleep(100 * time.Second)
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// go doSomething(&wg)
		go doSomething(&wg, i)
	}
	wg.Wait()
}
