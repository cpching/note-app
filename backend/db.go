package main

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB()  {
    var err error
    if db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/note_app"); err != nil {
        log.Fatal(err)
    } 
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
}


