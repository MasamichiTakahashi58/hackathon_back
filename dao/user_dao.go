package dao

import (
	"database/sql"
	"fmt"
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
    query := `SELECT id, email, username, display_name, bio, profile_image, header_image FROM users WHERE id = ?`
    row := db.DB.QueryRow(query, userID)

    var user model.User
    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.DisplayName, &user.Bio, &user.ProfileImage, &user.HeaderImage)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return &user, err
}

func GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, username, display_name, bio, profile_image, header_image FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.DisplayName, &user.Bio, &user.ProfileImage, &user.HeaderImage)
	if err == sql.ErrNoRows {
		return nil, nil // ユーザーが見つからない場合
	}
	return &user, err
}

func UpdateUser(user *model.User) error {
    query := "UPDATE users SET "
    params := []interface{}{}

    if user.Username != "" {
        query += "username = ?, "
        params = append(params, user.Username)
    }
    if user.DisplayName != "" {
        query += "display_name = ?, "
        params = append(params, user.DisplayName)
    }
    if user.Bio != nil {
        query += "bio = ?, "
        params = append(params, user.Bio)
    }
    if user.ProfileImage != nil {
        query += "profile_image = ?, "
        params = append(params, user.ProfileImage)
    }
    if user.HeaderImage != nil {
        query += "header_image = ?, "
        params = append(params, user.HeaderImage)
    }

    // クエリ末尾のカンマとスペースを削除
    query = query[:len(query)-2]
    query += " WHERE id = ?"
    params = append(params, user.ID)

    _, err := db.DB.Exec(query, params...)
    return err
}


func UpdateUserImage(userID int, column string, filePath string) error {
	query := fmt.Sprintf("UPDATE users SET %s = ? WHERE id = ?", column)
	_, err := db.DB.Exec(query, filePath, userID)
	return err
}
