package main

import (
	"fmt"
	"log"
)

func main() {
	x, y := 2, 3
	if x == 2 && y == 3 {
		doSomething()
	}

	if user, err := userName(); err == nil {
		fmt.Println(user.Name)
	} else {
		log.Println(err)
	}
}

func doSomething() {
	// do something
	return
}
