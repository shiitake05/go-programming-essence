package main

import "fmt"

type Value int

// 別名を付けた型にはメソッドを定義可能
func (v Value) Add(n Value) Value {
	return v + n
}

func main() {
	v := Value(1)
	v = v.Add(2)
	fmt.Println(v) // 3
}
