package store

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func Connection(dsn string) {

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	fmt.Println("connected")
}
