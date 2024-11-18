package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

func UpdateUserProfile(user *model.User) error {
    return dao.UpdateUserProfile(user)
}
