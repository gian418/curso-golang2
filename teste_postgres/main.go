package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.1.123"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "testego"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("Conexão ao banco de dados realizada com sucesso.")

	rows, err := db.Query("select * from movie")
	checkError(err)
	var id int
	var title string
	var year int
	var winner bool

	for rows.Next() {
		err = rows.Scan(&id, &title, &winner, &year)
		checkError(err)
		fmt.Println(id, " --- ", winner, " --- ", year, " --- ", title)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
