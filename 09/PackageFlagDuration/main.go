// この実装によってコマンドラインフラグに経過時間を渡すことが可能
// ./PackageFlagDuration -sleep 3s
package main

import (
	"flag"
	"time"
)

func main() {
	var sleep time.Duration
	flag.DurationVar(&sleep, "sleep", time.Second, "sleep time")
	flag.Parse()
	time.Sleep(sleep)
}
