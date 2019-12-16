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

// User struct
type User struct {
	ID        int
	Age       int
	Firstname string
	Lastname  string
	Email     string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
		SELECT * FROM users WHERE id = $1;
	`
	var user User
	// var id int
	// var email string

	row := db.QueryRow(sqlStatement, 4)
	err = row.Scan(&user.ID, &user.Age, &user.Firstname, &user.Lastname, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}
