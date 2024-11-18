package usecase

import (
    "hackathon_back/dao"
)

func ConfirmUser(email, code string) error {
    user, err := dao.GetTemporaryUser(email)
    if err != nil || user.VerificationCode != code {
        return err
    }
    return dao.ConfirmUser(email)
}
