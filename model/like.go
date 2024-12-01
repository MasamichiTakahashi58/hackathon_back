package model

type Like struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Username    string `json:"username"`       
    DisplayName string `json:"display_name"`  
    PostID    int    `json:"post_id"`
    CreatedAt string `json:"created_at"`
}
