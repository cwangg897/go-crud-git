package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/goprac")
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)
	fmt.Println("Connection Succesfully")

	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
