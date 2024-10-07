package main

import (
	"fmt"
)

// ランタイムpanicの場合はerror型の値が格納される
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T: %[1]s\n", e)
			// runtime.boundsError: runtime error: index out of range [2] with length 2
		}
	}()
	var a [2]int
	n := 2
	println(a[n])
}
