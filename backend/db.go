package main

import (
	"database/sql"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *sql.DB

func initDB() {
	var err error
	MYSQLDB_URI := os.Getenv("MYSQLDB_URI")

	if db, err = sql.Open("mysql", MYSQLDB_URI); err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Print("!")
		log.Fatal(err)
	}
}
