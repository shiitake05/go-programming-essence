// フレームワークを使用せずにルーティングを実装
// GETとPOSTも実装している

package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routeItem struct {
	method  string
	matcher *regexp.Regexp
	fnc     routerFunc
}

type router struct {
	items []routeItem
}

func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routeItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) POST(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routeItem{
		method:  http.MethodPost,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range rt.items {
		// リクエストメソッドの一致、リクエストURIがマッチ
		if v.method == r.Method && v.matcher.MatchString(r.RequestURI) {
			// パスパラメータの取得
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			param := make(routerParam)
			for i, name := range v.matcher.SubexpNames() {
				param[name] = match[i]
			}
			v.fnc(param, w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	rt := router{}
	// 先頭から見て/だけで終わる文字列
	rt.GET(`^/$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})
	// 先頭から見て/で始まり英数字だけが続いて終わる文字列
	rt.GET(`^/(?P<name>\w+)$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello: %v\n", p["name"])
	})
	// 先頭から見て/apiで終わる文字列
	rt.POST(`^/api$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/json")
		fmt.Fprint(w, `{"status": "OK"}`)
	})
	http.ListenAndServe(":8080", &rt)
}
