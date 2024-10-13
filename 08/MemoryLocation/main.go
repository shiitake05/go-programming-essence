// ベンチマークはパフォーマンス改善に役立つ
// 「go test -count 5 -benchmem -bench . 2>&1 | tee new.log」でベンチマークを計測
package MemoryLocation

import (
	"fmt"
)

// メモリアロケーションの効率が悪い例
// func makeSomething(n int) []string {
// 	var r []string
// 	for i := 0; i < n; i++ {
// 		r = append(r, fmt.Sprintf("%05d 何か", i))
// 	}
// 	return r
// }

// あらかじめ必要な長さのスライスを用意して代入するソースコード。スライスが都度、伸長しないので上のコードより効率が良い
// lopあたりのメモリ使用量が半分程度になる。また、lopあたりの実行速度も少ないことがわかる
// 「go install golang.org/x/perf/cmd/benchstat@latest」でbefore/afterのベンチマーク結果を比較できる
func makeSomething(n int) []string {
	r := make([]string, n, n)
	for i := 0; i < n; i++ {
		r[i] = fmt.Sprintf("%05d 何か", i)
	}
	return r
}