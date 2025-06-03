package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Failed to open database connection:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("Connected to MySQL database successfully")
}
