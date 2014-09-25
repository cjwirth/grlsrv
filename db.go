package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

var (
	Database *sql.DB
)

func init() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/totec")
	if err != nil {
		log.Fatal(err)
	}
	Database = db
	// it is long living, so do not close
	//	defer db.Close()
}
