package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "kchatman"
	dbname = "golang_db_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
		DELETE FROM users 
		WHERE id = $1
	`
	_, err = db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
}
