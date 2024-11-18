package controller

import (
    "net/http"
    "hackathon_back/usecase"
)

func VerifyUserHandler(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")
    code := r.FormValue("code")

    // メール認証処理
    if err := usecase.ConfirmUser(email, code); err != nil {
        http.Error(w, "Verification failed", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User registered successfully!"))
}
