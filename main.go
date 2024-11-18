package main

import (
    "net/http"
    "hackathon_back/controller"
    "hackathon_back/db"
    "log"
)

func main() {
    if err := db.ConnectDB(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.DB.Close()

    http.HandleFunc("/register", controller.RegisterUserHandler)
    http.HandleFunc("/verify", controller.VerifyUserHandler)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
