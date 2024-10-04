package main

import "fmt"

type User2 struct {
	Name string
	Age  int
}

func showUser(user2 *User2) {
	fmt.Println(user2.Name)
}

func main() {
	type User struct {
		Name string
		Age  int
	}

	// var user User
	// user.Name = "Bob"
	// user.Age = 18

	// JSON文字列を得る場合も上と同じように行える
	user := User{
		Name: "Bob",
		Age:  18,
	}

	fmt.Println(user)

	// コピーのオーバーヘッドをなくすのであればポインタを使う
	user2 := User2{
		Name: "Alice",
		Age:  20,
	}
	showUser(&user2)
}
