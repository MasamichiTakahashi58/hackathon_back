package dao

import (
    "hackathon_back/db"
    "hackathon_back/model"
)

func InsertTemporaryUser(user model.TemporaryUser) error {
    query := `INSERT INTO temporary_users (email, password, username, verification_code, expires_at) VALUES (?, ?, ?, ?, ?)`
    _, err := db.DB.Exec(query, user.Email, user.Password, user.Username, user.VerificationCode, user.ExpiresAt)
    return err
}

func GetTemporaryUser(email string) (*model.TemporaryUser, error) {
    var user model.TemporaryUser
    query := `SELECT email, verification_code, expires_at FROM temporary_users WHERE email = ?`
    err := db.DB.QueryRow(query, email).Scan(&user.Email, &user.VerificationCode, &user.ExpiresAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func ConfirmUser(email string) error {
    query := `INSERT INTO users (email, password, username) SELECT email, password, username FROM temporary_users WHERE email = ?`
    _, err := db.DB.Exec(query, email)
    if err != nil {
        return err
    }
    _, _ = db.DB.Exec("DELETE FROM temporary_users WHERE email = ?", email)
    return nil
}

