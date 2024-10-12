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