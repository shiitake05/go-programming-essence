// Goのテストスイートにはプロファイルを得る仕組みが用意されている
// メモリを大量に確保する重たいプログラム
// 長さの決まったスライスに文字列を追加し続ける処理
// package main

// import (
// 	"sync"
// )

// func heavyFunc(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	s := make([]string, 3)
// 	for i := 0; i < 1000000; i++ {
// 		s = append(s, "magical pandas")
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go heavyFunc(&wg)
// 	wg.Wait()
// }

// 上記のプログラムをプロファイルングを取るように変更する
// CPUプロファイルを得るには「./profile -cpuprofile cpu.prof」
// CPU負荷の高いものがわかる
// cumソートした結果を見るには「go tool pprof -cum -top cpu.prof」
// メモリプロファイルを得るには「./profile -memprofile mem.prof」
// メモリ負荷の高いものがわかる
// -pngフラグをつけることで呼び出しの相関がわかる
// CPUプロファイルの相関グラフを生成するには「go tool pprof -png cpu.prof > out.png」
// Webブラウザで相関やフレームグラフを表示するには「go tool pprof -http=:8080 cpu.prof」
// package main

// import (
// 	"flag"
// 	"log"
// 	"os"
// 	"runtime/pprof"
// 	"sync"
// )

// func heavyFunc(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	s := make([]string, 3)
// 	for i := 0; i < 1000000; i++ {
// 		s = append(s, "magical pandas")
// 	}
// }

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
// var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

// func main() {
// 	flag.Parse()
// 	if *cpuprofile != "" {
// 		f, err := os.Create(*cpuprofile)	// 処理本体の前にCPUプロファイルの開始
// 		if err != nil {
// 			log.Fatal("could not create CPU profile: ", err)
// 		}
// 		defer f.Close()
// 		if err := pprof.StartCPUProfile(f); err != nil {
// 			log.Fatal("could not start CPU profile: ", err)
// 		}
// 		defer pprof.StopCPUProfile()	// deferによる停止の予約
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go heavyFunc(&wg)
// 	wg.Wait()

// 	if *memprofile != "" {
// 		f, err := os.Create(*memprofile)	// 本体処理の後にヒープのプロファイル出力
// 		if err != nil {
// 			log.Fatal("could not create memory profile: ", err)
// 		}
// 		defer f.Close()
// 		// runtime.GC()
// 		if err := pprof.WriteHeapProfile(f); err != nil {
// 			log.Fatal("could not write memory profile: ", err)
// 		}
// 	}
// }

// 上記のプログラムをもう少し簡単にするために、new/http/pprofでWebサーバを起動する。
// また、プログラムがすぐに終了してしまわないように処理をループする
// これを「go run main.go」した後、別の端末から「curl -s http://localhost:6060/debug/pprof/profile > cpu.prof」でプロファイルを得ることができる
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func heavyFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 1000000; i++ {
		s = append(s, "magical pandas")
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		var wg sync.WaitGroup
		wg.Add(1)
		go heavyFunc(&wg)
		wg.Wait()
	}
}