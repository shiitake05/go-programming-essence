package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
)

// runではなくて良いが、処理の決まりとしてプログラム名を除いで引数文字列の配列と、終了コードを返すためにintの戻り値を取る点がある
func run(args []string) int {
	max := commandLine.Int("max", 255, "max value")
	name := commandLine.String("name", "", "my name")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
	}

	if *max > 999 {
		fmt.Fprintf(os.Stderr, "invalid max value: %d\n", *max)
		return 1
	}

	if *name == "" {
		fmt.Fprintln(os.Stderr, "name must be provided")
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
