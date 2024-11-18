package model

type Reply struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    PostID    int    `json:"post_id"`
    Content   string `json:"content"`
    ParentID  int    `json:"parent_id,omitempty"`
    CreatedAt string `json:"created_at"`
}
