package main

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)


var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/websitesate_user")
    if err != nil {
        log.Fatal(err)
    }
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
    log.Println("Database connected.")
}
