package model

type Post struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Content   string `json:"content"`
    ImageURL  string `json:"image_url,omitempty"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
