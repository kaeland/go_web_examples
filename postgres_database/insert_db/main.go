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
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	id := 0
	err = db.QueryRow(sqlStatement, 27, "kaeland1@gmail.com", "kaeland", "chatman").Scan(&id)
	err = db.QueryRow(sqlStatement, 30, "tim@gmail.com", "Tim", "James").Scan(&id)
	err = db.QueryRow(sqlStatement, 27, "will@gmail.com", "Will", "Thomas").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
