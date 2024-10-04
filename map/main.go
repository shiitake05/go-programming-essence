package main

import (
	"fmt"
	"sort"
)

func main() {
	// var m map[string]int

	// m = make(map[string]int)
	// m["John"] = 21
	// m["Bob"] = 18
	// m["Mark"] = 33

	// capを指定することもできる
	// m = make(map[string]int, 10)
	// m["John"] = 21
	// m["Bob"] = 18
	// m["Mark"] = 33

	// リテラルを使って初期値を代入することも可能
	m := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}
	fmt.Println(m)

	// mapから要素を削除するにはdelete関数を使う
	delete(m, "Bob")

	// mapのキーと値を列挙するためにはfor rangeを使う
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// ソートされたキーで列挙したい時はキーを取り出してソートする
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, m[k])
	}

	// mapは存在しないキーを指定すると、値の型のゼロ値が返る
	m0 := map[string]string{
		"foo": "bar",
	}
	fmt.Println(m0["zoo"])

	// キーが存在したかどうかを確認するためには、2つ目の戻り値を使う
	v, ok := m0["zoo"]
	if ok {
		// zooが存在すればokはtrueになる
		fmt.Println(v)
	}
}
