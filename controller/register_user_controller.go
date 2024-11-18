package controller

import (
    "encoding/json"
    "net/http"
    "hackathon_back/usecase"
    "hackathon_back/model"
    "fmt"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
    var user model.TemporaryUser
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // 仮ユーザー登録
    if err := usecase.RegisterTemporaryUser(user); err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    fmt.Printf("Verification code sent to %s\n", user.Email)
    w.WriteHeader(http.StatusCreated)
}
