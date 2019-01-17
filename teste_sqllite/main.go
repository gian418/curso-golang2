package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("teste2", "info2", "2019-01-14")
	checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)
	// // update
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("astaxieupdate", id)
	// checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Linhas afetadas: ", affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	rows.Close() //good habit to close

	// delete
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)

	//res, err = stmt.Exec(id)
	//checkErr(err)

	//affect, err = res.RowsAffected()
	//checkErr(err)

	//fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}