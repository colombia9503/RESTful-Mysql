package common

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

//DB Connection
func connectDB() {
	DB, err = sql.Open("mysql", "root@tcp(localhost:3306)/CPP")
	if err != nil {
		fmt.Println(err.Error())
	}

	//test connection
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
}
