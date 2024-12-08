package model

type Reply struct {
    ID        int    `json:"id"`
    UserID    int    `json:"user_id"`
    Username    string `json:"username"`       
    DisplayName string `json:"display_name"`  
    PostID    int    `json:"post_id"`
    Content   string `json:"content"`
    ParentID  *int    `json:"parent_id,omitempty"`
    CreatedAt string `json:"created_at"`
    RelationDepth int   `json:"relation_depth,omitempty"`
}
