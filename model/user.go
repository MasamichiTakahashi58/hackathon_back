package model

type User struct {
    ID           int    `json:"id"`
    Email        string `json:"email"`
    Username     string `json:"username"`
    DisplayName  string `json:"display_name"`
    ProfileImage string `json:"profile_image"`
    HeaderImage  string `json:"header_image"`
    Bio          string `json:"bio"`
    Location     string `json:"location"`
    Website      string `json:"website"`
    Birthdate    string `json:"birthdate"`
}
