package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "os"
)

var DB *sql.DB

func ConnectDB() error {
    dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("MYSQL_DATABASE")
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Failed to create database handle: %v", err)
        return err
    }
    if err = DB.Ping(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
        return err
    }
    log.Println("Database connection successful")
    return nil
}


