package main

import "fmt"

func main() {

	// 配列（固定長）
	var a [4]int
	a[0] = 1

	// スライス（可変長）
	// var x []int	// panicとなるため、makeで初期化する
	b := make([]int, 3)
	b[0] = 1
	b[1] = 2
	b[2] = 3

	// 初期値を代入することも可能
	c := []int{1, 2, 3}

	// 固定長2次元の配列
	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// 2次元のスライス
	arr2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println("-----------------------------")

	// スライスを伸長する
	d := []int{1, 2, 3}
	fmt.Println(d)
	d = append(d, 4, 5, 6)
	fmt.Println(d)
	fmt.Printf("d の長さは %d\n", len(d))

	e := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		e = append(e, i)
	}
	fmt.Println(e)

	// スライスの一部を取り出す
	fmt.Println(e[3:10]) // [3 4 5 6 7 8 9]

	// スライスの要素を削除する
	e2 := make([]int, 0, len(e)/2)
	for i := 0; i < len(e); i++ {
		if i%2 == 0 {
			e2 = append(e2, e[i])
		}
	}
	// e = e2
	fmt.Println(e2)

	// appendを使う方法
	n := 50
	e3 := append(e[:n], e[n+1:]...) // 添え字50を削除

	// copyを使う方法
	e4 := e[:n+copy(e[n:], e[n+1:])]

	fmt.Println(e3)
	fmt.Println(e4)
}
