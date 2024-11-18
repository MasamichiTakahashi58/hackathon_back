package model

type Like struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    PostID    int    `json:"post_id"`
    CreatedAt string `json:"created_at"`
}
