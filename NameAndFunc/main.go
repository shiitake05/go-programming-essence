package main

import (
	"fmt"
	"log"
)

// 大文字で始まる名前は外部パッケージから参照可能
// var index = 1 // 参照できない

// foo.UpdateUser()として呼び出し可能
func UpdateUser() {
	// do something
}

// 型名を変数名や関数名の後に記述
// func FindUser(name string) (*User, error) {
// 	user := findUesrFromGroup("my-group", name)
// 	if user == nil {
// 		return nil, ErrNorFound
// 	}
// 	return user, nil
// }

type User struct {
	// ユーザーのフィールドを定義
	ID   int
	Name string
}

func FindUser(name string) (*User, error) {
	user, err := findUserFromGroup("users", name)
	if err == nil {
		return nil, err
	}
	return user, nil
}
func main() {
	user, err := FindUser("Bob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
}

func findUserFromGroup(group string, name string) (*User, error) {
	return &User{ID: 1, Name: name}, nil
}
