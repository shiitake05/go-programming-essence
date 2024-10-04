package main

import (
	"fmt"
	"log"
)

func main() {
	x, y := 2, 3
	i := 1
	hit := 5
	if x == 2 && y == 3 {
		doSomething()
	}

	if user, err := userName(); err == nil {
		fmt.Println(user.Name)
	} else {
		log.Println(err)
	}

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	switch i {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	switch {
	case i > hit:
		fmt.Println("bigger than", hit)
	case i < hit:
		fmt.Println("less than", hit)
	case i == hit:
		fmt.Println("equal to", hit)
	}
}

func doSomething() {
	// do something
	return
}

func userName() (User, error) {
	return User{Name: "John"}, nil
}

type User struct {
	Name string
}
