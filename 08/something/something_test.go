// DoSomething関数の計測する
// 「go test -bench DoSomething」を実行
package something

import (
	"testing"
)

func BenchmarkDoSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSomething()
	}
}

// 「go test -bench .」で全てのベンチマークを実行
// 「go test -benchmem -bench DoSomething」でメモリロケーション量も計測できる