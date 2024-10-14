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
		err := db.QueryRow(`SELECT name, age FROM users WHERE id = $1`, id).Scan(&name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("name: %s, age: %d", name, age)
	}

	// 遅いクエリを途中でキャンセルすることも可能
	// ctx, cancel := context.WithCancel(context.Background(), 3*time.Second)
	// defer cancel()
	// rows, err := db.QueryContext(ctx, "SELECT name, age FROM users WHERE id < $1", id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
