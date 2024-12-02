package dao

import (
    "hackathon_back/db"
    "hackathon_back/model"
)

func CreatePost(post *model.Post) error {
    query := `INSERT INTO posts (user_id, content, image_url) VALUES (?, ?, ?)`
    _, err := db.DB.Exec(query, post.UserID, post.Content, post.ImageURL)
    return err
}


func GetPosts() ([]model.Post, error) {
    query := `
        SELECT 
            posts.id AS post_id,
            posts.user_id,
            posts.content,
            posts.image_url,
            posts.created_at,
            posts.updated_at,
            users.username,
            users.display_name
        FROM 
            posts
        JOIN 
            users ON posts.user_id = users.id
        ORDER BY 
            posts.created_at DESC
    `
    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []model.Post
    for rows.Next() {
        var post model.Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt, &post.Username, &post.DisplayName)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    if posts == nil {
        return []model.Post{}, nil
    }
    return posts, nil
}

func DeletePost(postID int) error {
    query := `DELETE FROM posts WHERE id = ?`
    _, err := db.DB.Exec(query, postID)
    return err
}
