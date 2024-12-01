package dao

import (
	"database/sql"
	"hackathon_back/db"
	"hackathon_back/model"
)

func CreateReply(reply *model.Reply) error {
    query := `INSERT INTO replies (user_id, post_id, content, parent_reply_id) VALUES (?, ?, ?, ?)`
    _, err := db.DB.Exec(query, reply.UserID, reply.PostID, reply.Content, reply.ParentID)
    return err
}

func GetRepliesByPost(postID int) ([]model.Reply, error) {
    query := `
        SELECT 
            r.id AS reply_id, 
            r.user_id, 
            u.username, 
            u.display_name, 
            r.post_id, 
            r.content, 
            r.parent_reply_id, 
            r.created_at
        FROM
            replies r
        JOIN
            users u ON r.user_id = u.id
        WHERE 
            r.post_id = ?
    `
    rows, err := db.DB.Query(query, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var replies []model.Reply
    for rows.Next() {
        var reply model.Reply
        var parentReplyID sql.NullInt64
        err := rows.Scan(&reply.ID, &reply.UserID, &reply.Username, &reply.DisplayName, &reply.PostID, &reply.Content, &parentReplyID, &reply.CreatedAt,
        )
        if err != nil {
            return nil, err
        }

        if parentReplyID.Valid {
            id := int(parentReplyID.Int64)
            reply.ParentID = &id
        } else {
            reply.ParentID = nil
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
