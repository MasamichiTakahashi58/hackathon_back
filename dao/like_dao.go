package dao

import (
    "hackathon_back/db"
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
