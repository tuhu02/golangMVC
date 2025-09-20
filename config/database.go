package config

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar-golang?parseTime=true")
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    log.Println("Database connected")
    DB = db
}