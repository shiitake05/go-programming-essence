//go:embedで文字列を扱う

package main

import (
	_ "embed"
	"fmt"
)

//go:embed message.txt
var message string

func main() {
	fmt.Println(message)
}
