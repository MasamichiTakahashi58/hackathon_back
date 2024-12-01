package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "hackathon_back/usecase"
    "hackathon_back/model"
    "log"
)

// 投稿の作成
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
    var post model.Post
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        log.Printf("Error creating post: %v", err)
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err = usecase.CreateNewPost(&post)
    if err != nil {
        log.Printf("Error creating post: %v", err)
        http.Error(w, "Failed to create post", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

// 投稿の取得
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
    posts, err := usecase.FetchAllPosts()
    if err != nil {
        log.Printf("Error fetching posts: %v", err)
        http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

// 投稿の削除
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
    postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
    if err != nil {
        log.Printf("Invalid post ID: %v", err)
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    err = usecase.RemovePost(postID)
    if err != nil {
        log.Printf("Error deleting post with ID %d: %v", postID, err)
        http.Error(w, "Failed to delete post", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
