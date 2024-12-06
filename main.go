package main

import (
	"hackathon_back/controller"
	"hackathon_back/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") 
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    println("!!###########################################################################################################")

    if err := godotenv.Load("./.env"); err != nil {
        log.Printf("Error loading .env file: %v", err)
    }

    if err := db.ConnectDB(); err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.DB.Close()

    // ユーザー
    http.HandleFunc("/users/create", controller.CreateUserHandler)
    http.HandleFunc("/users/update", controller.UpdateUserHandler)
    http.HandleFunc("/users/get", controller.GetUserHandler)
    http.HandleFunc("/users/email", controller.GetUserByEmailHandler)
    
    http.HandleFunc("/upload/user-image", controller.UploadUserImageHandler)

    // 投稿
    http.HandleFunc("/post/create", controller.CreatePostHandler)
    http.HandleFunc("/post/get", controller.GetPostsHandler)
    http.HandleFunc("/post/delete", controller.DeletePostHandler)

    // いいね
    http.HandleFunc("/like/add", controller.AddLikeHandler)
    http.HandleFunc("/like/remove", controller.RemoveLikeHandler)
    http.HandleFunc("/like/count", controller.CountLikesHandler)
    http.HandleFunc("/likes", controller.GetLikesByPostIDHandler)

    http.HandleFunc("/like/hasLiked", controller.HasUserLikedHandler)

    // リプライ
    http.HandleFunc("/reply/create", controller.CreateReplyHandler)
    http.HandleFunc("/reply/get", controller.GetRepliesHandler)
    http.HandleFunc("/reply/delete", controller.DeleteReplyHandler)

    // Gemini
    http.HandleFunc("/api/generate", controller.GenerateContentHandler)

    // ポート番号を環境変数 PORT から取得
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // デフォルトポート
    }

    // サーバー起動
    log.Printf("Server started on :%s", port)
    http.ListenAndServe(":"+port, enableCORS(http.DefaultServeMux))
}

