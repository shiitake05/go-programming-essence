// テストコード
package TestMain

import (
	"log"
	"os"
	"testing"
)

// テストの前処理と後処理
// 前処理と後処理の最中にエラーが発生した場合、os.Exitで異常終了したことを伝える
// 終了コード「0」以外で終了したい時は「log.Fatal」を使うと便利
// 例えば下記のような感じで
// func TestMain(m *testing.M) {
// 	if err := setup(); err != nil {
// 		log.Fatalf("failed to setup:", err)
// 	}

// 	ret := m.Run()

// 	if err := teardown(); err != nil {
// 		log.Fatalf("failed to teardown:", err)
// 	}
// 	os.Exit(ret)
// }

func TestMain(m *testing.M) {
	log.Println("Before")
	ret := m.Run()
	log.Println("After")
	os.Exit(ret)
}

func TestA(t *testing.T) {
	log.Println("TestA runnning")
}

func TestB(t *testing.T) {
	log.Println("TestB runnning")
}
