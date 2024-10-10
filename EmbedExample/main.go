// http://localhost:8989/ で画像を表示
package main

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 以下のパスがコンパイル時に読み取られ、プログラム実行時にバイト列として格納される(文字列(EmbedString)もディレクトリ(EmbedDirectory)も扱える)

//go:embed static/logo.png
var contents []byte

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/png", contents)
	})
	e.Logger.Fatal(e.Start(":8989"))
}
