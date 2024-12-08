package dao

import (
	"database/sql"
	"hackathon_back/db"
	"hackathon_back/model"
)

func CreateReply(reply *model.Reply) (int, error) {
    query := `
        INSERT INTO replies (user_id, post_id, content, parent_reply_id) 
        VALUES (?, ?, ?, ?)
    `
    result, err := db.DB.Exec(query, reply.UserID, reply.PostID, reply.Content, reply.ParentID)
    if err != nil {
        return 0, err
    }

    // 挿入されたリプライのIDを取得
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func AddReplyRelation(postID int, parentID *int, replyID int) error {
    var relationDepth int

    if parentID != nil {
        // 親リプライのrelation_depthを取得して+1
        query := `
            SELECT relation_depth
            FROM reply_relations
            WHERE reply_id = ?
        `
        row := db.DB.QueryRow(query, *parentID)
        if err := row.Scan(&relationDepth); err != nil {
            return err
        }
        relationDepth++ // 親リプライの深さ +1
    } else {
        // 親リプライがない場合（ポスト直下）
        relationDepth = 1
    }

    // 新しいリレーションを挿入
    query := `
        INSERT INTO reply_relations (post_id, parent_reply_id, reply_id, relation_depth)
        VALUES (?, ?, ?, ?)
    `
    _, err := db.DB.Exec(query, postID, parentID, replyID, relationDepth)
    return err
}



func GetRepliesByPostWithRelations(postID int) ([]model.Reply, error) {
    query := `
        SELECT 
            r.id, r.user_id, u.username, u.display_name,
            r.post_id, r.content, r.parent_reply_id,
            rr.relation_depth, r.created_at
        FROM replies r
        JOIN reply_relations rr ON r.id = rr.reply_id
        JOIN users u ON r.user_id = u.id
        WHERE rr.post_id = ?
        ORDER BY rr.relation_depth, rr.parent_reply_id
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
        err := rows.Scan(
            &reply.ID, &reply.UserID, &reply.Username, &reply.DisplayName,
            &reply.PostID, &reply.Content, &parentReplyID,
            &reply.RelationDepth, &reply.CreatedAt,
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

func DeleteReplyRelation(replyID int) error {
    query := `DELETE FROM reply_relations WHERE reply_id = ? OR parent_reply_id = ?`
    _, err := db.DB.Exec(query, replyID, replyID)
    return err
}

// ポストIDでリプライを削除
func DeleteRepliesByPostID(postID int) error {
    // リレーションから削除
    relationQuery := `DELETE FROM reply_relations WHERE post_id = ?`
    _, err := db.DB.Exec(relationQuery, postID)
    if err != nil {
        return err
    }

    // リプライ本体を削除
    replyQuery := `DELETE FROM replies WHERE post_id = ?`
    _, err = db.DB.Exec(replyQuery, postID)
    return err
}
