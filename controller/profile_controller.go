package controller

import (
    "encoding/json"
    "net/http"
    "hackathon_back/usecase"
    "hackathon_back/model"
)

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err = usecase.UpdateUserProfile(&user)
    if err != nil {
        http.Error(w, "Failed to update profile", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
