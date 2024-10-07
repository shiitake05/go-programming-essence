package main

import (
	"fmt"
)

// 文字列を引数にしてpanic関数を呼び出した場合は、recoverから返される値にxstring型の値が格納される
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T: %[1]s\n", e)
			// string: my error
		}
	}()
	panic("my error")
}
