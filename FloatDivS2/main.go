// 引数を割り算し浮動小数点の値として結果を表示するプログラム
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor must not be zero")
	}
	return a / b, nil
}

func exitIf(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", message, err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Number of arguments must be 2")
		os.Exit(1)
	}
	a, err := strconv.ParseFloat(os.Args[1], 64)
	exitIf(err, "First argument must be a float value")
	b, err := strconv.ParseFloat(os.Args[2], 64)
	exitIf(err, "Second argument must be a float value")
	result, err := divide(a, b)
	exitIf(err, "Invalid argument")
	fmt.Println(result)
}
