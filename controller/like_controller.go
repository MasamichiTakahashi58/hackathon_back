package controller

import (
    "net/http"
    "strconv"
    "hackathon_back/dao"
)

func AddLikeHandler(w http.ResponseWriter, r *http.Request) {
    userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
    postID, _ := strconv.Atoi(r.URL.Query().Get("post_id"))

    err := dao.AddLike(userID, postID)
    if err != nil {
        http.Error(w, "Failed to add like", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func RemoveLikeHandler(w http.ResponseWriter, r *http.Request) {
    userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
    postID, _ := strconv.Atoi(r.URL.Query().Get("post_id"))

    err := dao.RemoveLike(userID, postID)
    if err != nil {
        http.Error(w, "Failed to remove like", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func CountLikesHandler(w http.ResponseWriter, r *http.Request) {
    postID, _ := strconv.Atoi(r.URL.Query().Get("post_id"))

    count, err := dao.CountLikes(postID)
    if err != nil {
        http.Error(w, "Failed to count likes", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(strconv.Itoa(count)))
}
