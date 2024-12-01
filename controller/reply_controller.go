package controller

import (
    "encoding/json"
    "net/http"
    "strconv"
    "hackathon_back/usecase"
    "hackathon_back/model"
    "log"
)

func CreateReplyHandler(w http.ResponseWriter, r *http.Request) {
    var reply model.Reply
    err := json.NewDecoder(r.Body).Decode(&reply)
    if err != nil {
        log.Printf("Error decoding request body: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if reply.Content == "" {
        http.Error(w, "Content is required", http.StatusBadRequest)
        return
    }

    err = usecase.CreateNewReply(&reply)
    if err != nil {
        log.Printf("Error creating reply in usecase: %v", err)
        http.Error(w, "Failed to create reply", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(reply)
}


func GetRepliesHandler(w http.ResponseWriter, r *http.Request) {
    postIDStr := r.URL.Query().Get("post_id")
    if postIDStr == "" {
        http.Error(w, "post_id is required", http.StatusBadRequest)
        return
    }

    postID, err := strconv.Atoi(postIDStr)
    if err != nil {
        http.Error(w, "Invalid post_id", http.StatusBadRequest)
        return
    }

    replies, err := usecase.FetchRepliesByPost(postID)
    if err != nil {
        http.Error(w, "Failed to fetch replies", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
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
