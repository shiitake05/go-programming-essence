// 複数のクエリや実行を扱いたい場合のために、ステートメントを先に作成しておくことができる
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var name string
	var id int
	var age int

	db, err := sql.Open("postgres", "postgres://postgres:postgres@dbserver/database")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`SELECT name, age FROM users WHERE id < $1`)
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(id).Scan(&name, &age)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("name: %v, age: %v", name, age)
}
