package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"
)

func RegisterTemporaryUser(user model.TemporaryUser) error {
    return dao.InsertTemporaryUser(user)
}
