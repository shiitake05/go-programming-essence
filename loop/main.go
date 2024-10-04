package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	i := 0
	obj := map[string]string{"key1": "value1", "key2": "value2"}
	arr := [3]int{1, 2, 3}

	// for形式
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// while形式
	for i < len(s) {
		fmt.Println(s[i])
		i++
	}

	// 無限ループ
	// for {
	// 	fmt.Println(s[i])
	// }

	// map
	for k, v := range obj {
		fmt.Println(k, v)
	}

	// slice
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// Labeled Breakという記法によって多重ループの内側から外側に抜けることが可能
finished:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if check(i, j) {
				break finished
			}
		}
	}

	// channelの読み取りをforで行うことも可能
	// for v:= range ch {
	// 	fmt.Println(v)
	// }
}

func check(i, j int) bool {
	return i == 99 && j == 50
}
