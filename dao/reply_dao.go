package dao

import (
    "hackathon_back/db"
    "hackathon_back/model"
)

func CreateReply(reply *model.Reply) error {
    query := `INSERT INTO replies (user_id, post_id, content, parent_id) VALUES (?, ?, ?, ?)`
    _, err := db.DB.Exec(query, reply.UserID, reply.PostID, reply.Content, reply.ParentID)
    return err
}

func GetRepliesByPost(postID int) ([]model.Reply, error) {
    rows, err := db.DB.Query(`SELECT id, user_id, post_id, content, parent_id, created_at FROM replies WHERE post_id = ?`, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var replies []model.Reply
    for rows.Next() {
        var reply model.Reply
        err := rows.Scan(&reply.ID, &reply.UserID, &reply.PostID, &reply.Content, &reply.ParentID, &reply.CreatedAt)
        if err != nil {
            return nil, err
        }
        replies = append(replies, reply)
    }
    return replies, nil
}

func DeleteReply(replyID int) error {
    query := `DELETE FROM replies WHERE id = ?`
    _, err := db.DB.Exec(query, replyID)
    return err
}
