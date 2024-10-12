// テストコード
package hsd

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Before")
	ret := m.Run()
	log.Println("After")
	os.Exit(ret)
}

// 各テストで-shortのときスキップしたい場合はt.SkipNow()を使う
func TestA(t *testing.T) {
	// -shortをつけて実行した際にtesting.Short()がtrueになる
	if testing.Short() {
		t.SkipNow()
	}
	log.Println("TestA runnning")
}

func TestB(t *testing.T) {
	log.Println("TestB runnning")
}
