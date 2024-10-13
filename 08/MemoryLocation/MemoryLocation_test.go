// ベンチマークコード
package MemoryLocation

import (
	"testing"
)

func BenchmarkMemoryLocation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = makeSomething(1000)
	}
}