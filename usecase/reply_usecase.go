package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

func CreateNewReply(reply *model.Reply) error {
    return dao.CreateReply(reply)
}

func FetchRepliesByPost(postID int) ([]model.Reply, error) {
    return dao.GetRepliesByPost(postID)
}

func RemoveReply(replyID int) error {
    return dao.DeleteReply(replyID)
}
