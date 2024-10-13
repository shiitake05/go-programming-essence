// 1文字の幅がアスキー文字よりも大きな文字を扱う場合に便利
// 端末に表示する文字列にアスキー文字で下線を引きたい場合に使える
package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

func main() {
	s := "Go言語でCLIアプリケーション作成"
	fmt.Println(s)
	width := runewidth.StringWidth(s)
	fmt.Println(strings.Repeat("~", width))

	// 指定幅での折り返しは以下のように記述する
	fmt.Println(runewidth.Wrap(s, 11))
}
