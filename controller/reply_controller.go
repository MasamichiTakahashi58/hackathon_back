package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "hackathon_back/usecase"
    "hackathon_back/model"
)

func CreateReplyHandler(w http.ResponseWriter, r *http.Request) {
    var reply model.Reply
    err := json.NewDecoder(r.Body).Decode(&reply)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    
    if reply.Content == "" {
        http.Error(w, "Content is required", http.StatusBadRequest)
        return
    }

    err = usecase.CreateNewReply(&reply)
    if err != nil {
        http.Error(w, "Failed to create reply", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(reply)
}

func GetRepliesHandler(w http.ResponseWriter, r *http.Request) {
    postID, _ := strconv.Atoi(r.URL.Query().Get("post_id"))
    replies, err := usecase.FetchRepliesByPost(postID)
    if err != nil {
        http.Error(w, "Failed to fetch replies", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(replies)
}

func DeleteReplyHandler(w http.ResponseWriter, r *http.Request) {
    replyID, _ := strconv.Atoi(r.URL.Query().Get("reply_id"))
    err := usecase.RemoveReply(replyID)
    if err != nil {
        http.Error(w, "Failed to delete reply", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
