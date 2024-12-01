package controller

import (
	"encoding/json"
	"hackathon_back/dao"
	"hackathon_back/usecase"
	"log"
	"net/http"
	"strconv"
)

func AddLikeHandler(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        PostID int `json:"post_id"`
        UserID int `json:"user_id"`
    }

    // リクエストボディをデコード
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    log.Print(r.Body)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    log.Println(requestBody.UserID, requestBody.PostID)

    err = dao.AddLike(requestBody.UserID, requestBody.PostID)
    if err != nil {
        http.Error(w, "Failed to add like", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func RemoveLikeHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    err = dao.RemoveLike(userID, postID)
    if err != nil {
        http.Error(w, "Failed to remove like", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func CountLikesHandler(w http.ResponseWriter, r *http.Request) {
    postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    count, err := dao.CountLikes(postID)
    if err != nil {
        http.Error(w, "Failed to count likes", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(strconv.Itoa(count)))
}

func GetLikesByPostIDHandler(w http.ResponseWriter, r *http.Request) {
    postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    likes, err := usecase.FetchLikesByPostID(postID)
    if err != nil {
        http.Error(w, "Failed to fetch likes", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(likes)
}

func HasUserLikedHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    liked, err := dao.HasUserLiked(userID, postID)
    if err != nil {
        http.Error(w, "Failed to check like status", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]bool{"liked": liked})
}