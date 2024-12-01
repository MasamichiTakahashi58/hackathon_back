package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

func AddLike(userID, postID int) error {
    //like check
    liked, err := dao.HasUserLiked(userID, postID)
    if err != nil {
        return err
    }

    if liked {
        return nil 
    }

    return dao.AddLike(userID, postID)
}

func RemoveLike(userID, postID int) error {
    return dao.RemoveLike(userID, postID)
}

func CountLikes(postID int) (int, error) {
    return dao.CountLikes(postID)
}

func FetchLikesByPostID(postID int) ([]model.Like, error) {
    return dao.GetLikesByPostID(postID)
}