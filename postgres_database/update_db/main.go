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
		UPDATE users 
		SET first_name = $2, last_name = $3 
		WHERE id = $1 
		RETURNING id, email;
	`
	var email string
	var id int
	err = db.QueryRow(sqlStatement, 1, "James", "Thomas").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)
}
