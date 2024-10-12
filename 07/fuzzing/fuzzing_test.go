// fuzzingテストと開発者が準備したテストの入力データではなく、開発者が予期しないであろうランダムで無効なデータを用いてテストする方法
package main

import (
	"testing"
)

func FuzzDoSomething(f *testing.F) {
	f.Add("test&&&")
	f.Fuzz(func(t *testing.T, data string) {
		doSomething(data)
	})
}
