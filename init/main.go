package main

// importも同様
// import (
// 	_ "my-init/init"
// )

// Goでデータベースを利用する場合はdatabase/sqlパッケージとデータベースドライバをimportする
// そのため、以下のようにドライバをブランクimportして利用する
// import (
// 	"database/sql"
// 	_ "github.com/lib/pq"
// )

var name = "John"

// init関数はmain関数よりも先に実行される
func init() {
	println("Hi! " + name)
	// sql.Register("postgres", &Driver{})
}

func main() {
	println("Hello! " + name)
}
