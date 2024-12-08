package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

func CreateNewReply(reply *model.Reply) error {
    replyID, err := dao.CreateReply(reply)
    if err != nil {
        return err
    }

    // リプライIDを設定
    reply.ID = replyID

    // リレーションテーブルを更新
    return dao.AddReplyRelation(reply.PostID, reply.ParentID, replyID)
}

func FetchRepliesByPost(postID int) ([]model.Reply, error) {
    return dao.GetRepliesByPostWithRelations(postID)
}

func RemoveReply(replyID int) error {
    if err := dao.DeleteReply(replyID); err != nil {
        return err
    }

    // リレーションテーブルの関連データを削除
    return dao.DeleteReplyRelation(replyID)
}

func RemoveRepliesByPostID(postID int) error {
    return dao.DeleteRepliesByPostID(postID)
}

