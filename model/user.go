package model

import "time"

// 本登録
type User struct {
    ID           int    `json:"id"`
    Email        string `json:"email"`
    Password     string `json:"-"`
    Username     string `json:"username"`
    DisplayName  string `json:"display_name"`
    ProfileImage string `json:"profile_image"`
    HeaderImage  string `json:"header_image"`
    Bio          string `json:"bio"`
    Location     string `json:"location"`
    Website      string `json:"website"`
    Birthdate    string `json:"birthdate"`
    CreatedAt    string `json:"created_at"`
}

// 仮登録
type TemporaryUser struct {
    Email            string    `json:"email"`
    Password         string    `json:"password"`
    Username         string    `json:"username"`
    VerificationCode string    `json:"verification_code"`
    ExpiresAt        time.Time `json:"expires_at"`
}
