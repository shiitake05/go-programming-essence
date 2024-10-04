package main

import (
	"fmt"
)

type Value int

func (v *Value) Add(n Value) {
	*v += n // レシーバをデリファレンスして代入
}

// func (v Value) Add(n Value) {
// 	v += n // vはコピーなので元の値には影響しない
// }

// エスケープ解析によって、スタックからヒープに切り替わるためクラッシュしない
// func Bob() *User {
// 	// 関数内の変数userのポインタを関数外へ返す
// 	user := User{
// 		Name: "Bob",
// 		Age:  18,
// 	}
// 	return &user
// }

func main() {
	v := 1
	p := &v
	*p = 2
	fmt.Println(v) // 2

	v2 := Value(1)
	v2.Add(2)
	fmt.Println(v2) // 3
}
