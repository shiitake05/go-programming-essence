// http://localhost:3000/ でHello, World!
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
)

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *AppContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {
	router := web.New(AppContext{}). // ルータを作成
						Middleware(web.LoggerMiddleware).        // 同梱されたミドルウェアを使用
						Middleware(web.ShowErrorsMiddleware).    // ...
						Middleware((*AppContext).SetHelloCount). // 自身のミドルウェア
						Get("/", (*AppContext).SayHello)         // ルーティングを追加

	http.ListenAndServe("localhost:3000", router) // サーバーを起動
}
