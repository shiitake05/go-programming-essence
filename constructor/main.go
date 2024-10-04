package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Goにはコンストラクタがないので、以下のように構造体の初期化を行う関数を作成することで代用する
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func main() {
	user := NewUser("Bob", 18)
	fmt.Println(user.Name, user.Age)
}
