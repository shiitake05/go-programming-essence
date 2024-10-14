// INSERTやDELETEを実行する場合に使用する
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	var name string
	var id int

	db, err := sql.Open("postgres", "postgres://postgres:postgres@dbserver/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Exec(`UPDATE users SET name = $1 WHERE id = $2`, name, id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(rows)
}
