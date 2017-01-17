package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func InitDatabase() {
	db, err := sql.Open("mysql", "root:noibuongactro@/test")
	if err != nil {
		fmt.Println("Couldn't initialize MySQL")
		fmt.Println(err)
	} else {
		fmt.Println("Initialize MySQL successfully")
	}
	Database = db
}
