package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

// ユーザーの作成
func CreateUser(user *model.User) error {
	return dao.CreateUser(user)
}

// ユーザーの取得
func GetUserByID(userID int) (*model.User, error) {
    return dao.GetUserByID(userID)
}

// ユーザーの更新
func UpdateUser(user *model.User) error {
    return dao.UpdateUser(user)
}

