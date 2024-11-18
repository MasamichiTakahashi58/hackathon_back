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
    rows, err := db.DB.Query(`SELECT id, user_id, content, image_url, created_at, updated_at FROM posts`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []model.Post
    for rows.Next() {
        var post model.Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func DeletePost(postID int) error {
    query := `DELETE FROM posts WHERE id = ?`
    _, err := db.DB.Exec(query, postID)
    return err
}
