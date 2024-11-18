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
        log.Fatal("Failed to connect to DB:", err)
        return err
    }
    return DB.Ping()
}
