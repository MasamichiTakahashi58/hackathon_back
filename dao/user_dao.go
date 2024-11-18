package dao

import (
    "hackathon_back/db"
    "hackathon_back/model"
)

func UpdateUserProfile(user *model.User) error {
    query := `UPDATE users SET display_name=?, profile_image=?, header_image=?, bio=?, location=?, website=?, birthdate=? WHERE email=?`
    _, err := db.DB.Exec(query, user.DisplayName, user.ProfileImage, user.HeaderImage, user.Bio, user.Location, user.Website, user.Birthdate, user.Email)
    return err
}


