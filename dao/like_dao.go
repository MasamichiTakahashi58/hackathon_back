package dao

import (
    "hackathon_back/db"
    "hackathon_back/model"
)

func AddLike(userID, postID int) error {
    query := `INSERT INTO likes (user_id, post_id) VALUES (?, ?)`
    _, err := db.DB.Exec(query, userID, postID)
    return err
}

func RemoveLike(userID, postID int) error {
    query := `DELETE FROM likes WHERE user_id = ? AND post_id = ?`
    _, err := db.DB.Exec(query, userID, postID)
    return err
}

func CountLikes(postID int) (int, error) {
    var count int
    query := `SELECT COUNT(*) FROM likes WHERE post_id = ?`
    err := db.DB.QueryRow(query, postID).Scan(&count)
    return count, err
}

func HasUserLiked(userID, postID int) (bool, error) {
    var exists bool
    query := `SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = ? AND post_id = ?)`
    err := db.DB.QueryRow(query, userID, postID).Scan(&exists)
    return exists, err
}

func GetLikesByPostID(postID int) ([]model.Like, error) {
    query := `
        SELECT 
            likes.id AS like_id
            likes.user_id,
            users.username,
            users.display_name,
            likes.post_id,
            likes.created_at
        FROM 
            likes
        JOIN 
            users ON likes.user_id = users.id
        WHERE 
            likes.post_id = ?
    `
    rows, err := db.DB.Query(query, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var likes []model.Like
    for rows.Next() {
        var like model.Like
        err := rows.Scan(&like.ID, &like.UserID, &like.Username, &like.DisplayName, &like.PostID, &like.CreatedAt)
        if err != nil {
            return nil, err
        }
        likes = append(likes, like)
    }
    return likes, nil
}