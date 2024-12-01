package model

type Post struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Username    string `json:"username"`       
    DisplayName string `json:"display_name"`  
    Content   string `json:"content"`
    ImageURL  string `json:"image_url,omitempty"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
