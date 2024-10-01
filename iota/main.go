package main

import "fmt"

type Fruit int
type Animal int

func main() {
	// Goには列挙型がない
	// 代わりに定数とiotaを使う
	// const (
	// 	Apple = iota
	// 	Orange
	// 	Banana
	// )
	// println(Apple, Orange, Banana) // 0 1 2

	// const (
	// 	Apple = iota + iota
	// 	Orange
	// 	Banana
	// )
	// println(Apple, Orange, Banana) // 0 2 4

	fmt.Println(Apple2)  // 0 + 0 = 0
	fmt.Println(Orange2) // 1 + 1 = 2
	fmt.Println(Banana2) // 2 + 3 = 5

	var fruit Fruit = Apple3
	fmt.Println(fruit) // 0
	// Fruit = Elephant	// コンパイルエラー
	fmt.Println(fruit) // 0
}

const (
	Apple2 = iota + iota
	Orange2
	Banana2 = iota + 3
)

const (
	Apple3  Fruit = iota // Fruit(0)
	Orange3              // Fruit(1)
	Banana3              // Fruit(2)
)

const (
	Monkey   Animal = iota // Animal(0)
	Elephant               // Animal(1)
	Pig                    // Animal(2)
)
