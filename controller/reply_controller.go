package controller

import (
    "encoding/json"
    "hackathon_back/model"
    "hackathon_back/usecase"
    "log"
    "net/http"
    "strconv"
)

// リプライ作成
func CreateReplyHandler(w http.ResponseWriter, r *http.Request) {
    var reply model.Reply

    // リクエストボディをデコード
    if err := json.NewDecoder(r.Body).Decode(&reply); err != nil {
        log.Printf("Error decoding request body: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // 必須フィールドのバリデーション
    if reply.Content == "" || reply.UserID == 0 || reply.PostID == 0 {
        log.Println("Missing required fields: content, user_id, or post_id")
        http.Error(w, "Content, user_id, and post_id are required", http.StatusBadRequest)
        return
    }

    // リプライ作成
    if err := usecase.CreateNewReply(&reply); err != nil {
        log.Printf("Error creating reply in usecase: %v", err)
        http.Error(w, "Failed to create reply", http.StatusInternalServerError)
        return
    }

    // 成功応答
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(reply); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}

// 指定ポストIDのリプライ取得
func GetRepliesHandler(w http.ResponseWriter, r *http.Request) {
    postIDStr := r.URL.Query().Get("post_id")
    if postIDStr == "" {
        log.Println("post_id is missing in query parameters")
        http.Error(w, "post_id is required", http.StatusBadRequest)
        return
    }

    postID, err := strconv.Atoi(postIDStr)
    if err != nil {
        log.Printf("Invalid post_id: %v", err)
        http.Error(w, "Invalid post_id", http.StatusBadRequest)
        return
    }

    // リプライ取得
    replies, err := usecase.FetchRepliesByPost(postID)
    if err != nil {
        log.Printf("Error fetching replies for post_id %d: %v", postID, err)
        http.Error(w, "Failed to fetch replies", http.StatusInternalServerError)
        return
    }

    // 成功応答
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(replies); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}

// リプライ削除
func DeleteReplyHandler(w http.ResponseWriter, r *http.Request) {
    replyIDStr := r.URL.Query().Get("reply_id")
    if replyIDStr == "" {
        log.Println("reply_id is missing in query parameters")
        http.Error(w, "reply_id is required", http.StatusBadRequest)
        return
    }

    replyID, err := strconv.Atoi(replyIDStr)
    if err != nil {
        log.Printf("Invalid reply_id: %v", err)
        http.Error(w, "Invalid reply_id", http.StatusBadRequest)
        return
    }

    // リプライ削除
    if err := usecase.RemoveReply(replyID); err != nil {
        log.Printf("Error deleting reply with ID %d: %v", replyID, err)
        http.Error(w, "Failed to delete reply", http.StatusInternalServerError)
        return
    }

    // 成功応答
    w.WriteHeader(http.StatusOK)
    if _, err := w.Write([]byte("Reply deleted successfully")); err != nil {
        log.Printf("Error writing response: %v", err)
    }
}
