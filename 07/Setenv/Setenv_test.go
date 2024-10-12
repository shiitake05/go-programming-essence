// Setenv()は環境変数に依存した処理をテストする際に便利
// テスト関数の終了とともに、環境変数が元に戻される
package main

import (
	"testing"
)

func doSomething() error {
	return nil
}

func TestCreateProfile(t *testing.T) {
	t.Setenv("DATABASE_URL", "localhost")
	err := doSomething()
	if err != nil {
		t.Fatalf("cannot do something: %v", err)
	}
}
