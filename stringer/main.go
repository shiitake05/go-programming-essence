package main

import (
	"fmt"
)

// Stringerインターフェースでfmtパッケージの関数に与えた引数を明示的な文字列として出力
type Stringer interface {
	String() string
}

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q, AGE=%d", f.Name, f.Age)
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
	}

	fmt.Printf("%v\n", f) // NAME="John", AGE=20
}
