package main

import (
    "log"
    "net/http"
    "hackathon_back/controller"
    "hackathon_back/db"
    "github.com/joho/godotenv" 
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    if err := db.ConnectDB(); err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.DB.Close()

    // プロフィール
    http.HandleFunc("/profile/update", controller.UpdateProfileHandler)

    // 投稿
    http.HandleFunc("/post/create", controller.CreatePostHandler)
    http.HandleFunc("/post/get", controller.GetPostsHandler)
    http.HandleFunc("/post/delete", controller.DeletePostHandler)

    // いいね
    http.HandleFunc("/like/add", controller.AddLikeHandler)
    http.HandleFunc("/like/remove", controller.RemoveLikeHandler)
    http.HandleFunc("/like/count", controller.CountLikesHandler)

    // リプライ
    http.HandleFunc("/reply/create", controller.CreateReplyHandler)
    http.HandleFunc("/reply/get", controller.GetRepliesHandler)
    http.HandleFunc("/reply/delete", controller.DeleteReplyHandler)

    // サーバーの起動
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
