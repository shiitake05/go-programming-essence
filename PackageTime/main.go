// 時間を扱うためのパッケージ
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339)) // 現在時刻をRFC3339形式で出力
}
