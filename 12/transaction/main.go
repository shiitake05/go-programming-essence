// トランザクションを扱い、複数の更新を行ったり、Commitしたり、Rollbackしたりすることができる
package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@dbserver/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	result, err := tx.Exec(`INSERT INTO users (name, age) VALUES ($1, $2)`, "Fred", 41)
	if err != nil {
		log.Fatal(err)
	}
	if affected, err := result.RowsAffected(); err != nil {
		log.Fatal(err)
	} else if affected == 0 {
		log.Fatal(errors.New("no rows affected"))
	}
	tx.Commit()
}
