package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func doSomething(dir string) error {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir) // 関数の終了時にディレクトリを削除する

	f, err := os.Create(filepath.Join(dir, "data.txt"))
	if err != nil {
		return err
	}
	defer f.Close() // 関数の終了時にファイルを閉じる

	_, err = f.Write([]byte("Hello, World!"))
	if err != nil {
		return err
	}
	return nil
}

// deferには無名関数を渡すことも可能
func doSomething2() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()

	n = 2
}

// 上の処理と同様
func doSomething3() {
	var n = 1
	n = 2
	func() {
		fmt.Println(n)
	}()
}

// deferに通常の関数を渡す
func doSomething4() {
	var n = 1
	defer fmt.Println(n)

	n = 2
}

// 上の処理と同様
func doSomething5() {
	// var n = 1
	// n = 2
	fmt.Println(1)
}

// deferは後処理を記述できる
func main() {
	f, err := os.Open("data.txt") // ファイルを開く
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // 関数の終了時に必ず実行される

	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))

	//fは上書きされるが、deferで正しくファイルが閉じられる
	f, err = os.Open("data01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f, err = os.Open("data02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// deferは指定された順とは逆順に実行される
	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

	doSomething("temp")

	// deferは関数スコープで実行されるため、forループ内でdeferを使うと、ループが終了するまで実行されない
	for i := 0; i < 100; i++ {
		f, err := os.Open("data.txt")
		if err != nil {
			return
		}
		defer f.Close()
	}

	doSomething2()
	doSomething3()
	doSomething4()
	doSomething5()
}
