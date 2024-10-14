// PostgreSQLを扱うQuery
// SELECTを実行する場合にQueryでステートメントを作成し、Nwxtを呼び出してレコードをフェッチ、Scanで値を抽出する流れ
// フェッチの件数が1軒の時はQueryRow使用する
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@dbserver/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	id := 3
	rows, err := db.Query("SELECT name, age FROM users WHERE id < $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("name: %s, age: %d", name, age)
	}
}
