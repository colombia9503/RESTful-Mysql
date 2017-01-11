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
	//defer DB.Close()

	//test connection
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	/**
	//table creation (if not exist)
	stmt, err := DB.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Person Table successfully migrated....")
	}*/

}
