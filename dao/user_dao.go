package dao

import (
    "database/sql"
    "hackathon_back/db"
    "hackathon_back/model"
)

func CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (email, username, display_name, profile_image, header_image, bio)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, user.Email, user.Username, user.DisplayName, user.ProfileImage, user.HeaderImage, user.Bio)
	return err
}

func GetUserByID(userID int) (*model.User, error) {
    query := `SELECT id, email, username, display_name, bio, profile_image FROM users WHERE id = ?`
    row := db.DB.QueryRow(query, userID)

    var user model.User
    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.DisplayName, &user.Bio, &user.ProfileImage)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &user, err
}

func UpdateUser(user *model.User) error {
    query := `
        UPDATE users 
        SET username = ?, display_name = ?, bio = ?, profile_image = ?
        WHERE id = ?
    `
    _, err := db.DB.Exec(query, user.Username, user.DisplayName, user.Bio, user.ProfileImage, user.ID)
    return err
}

