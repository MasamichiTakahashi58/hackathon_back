package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

func CreateNewPost(post *model.Post) error {
    return dao.CreatePost(post)
}

func FetchAllPosts() ([]model.Post, error) {
    return dao.GetPosts()
}

func RemovePost(postID int) error {
    return dao.DeletePost(postID)
}
