// Durationで経過時間を保持できる
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	d, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d) // 3s

	d, err = time.ParseDuration("4m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d) // 4m0s

	d, err = time.ParseDuration("5h")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d) // 5h0m0s

	// 四則演算も可能
	d, err = time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d * 3) // 9s

	//Duration型の定数はtime.Nanosecondを「１」として以下のようになっている
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
}
