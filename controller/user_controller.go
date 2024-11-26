package controller

import (
	"encoding/json"
	"log"
	"net/http"
    "strconv"
	"hackathon_back/model"
	"hackathon_back/usecase"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// リクエストボディのデコード
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("JSON decode error: %v", err)
		return
	}

	// 必須フィールドのチェック
	if user.Username == "" || user.DisplayName == "" || user.Email == "" {
		http.Error(w, "Email, Username, and Display Name are required", http.StatusBadRequest)
		return
	}

	// 空フィールドを処理
	user.ProfileImage = normalizeString(user.ProfileImage)
	user.HeaderImage = normalizeString(user.HeaderImage)
	user.Bio = normalizeString(user.Bio)

	
	// ユーザー作成
	if err := usecase.CreateUser(&user); err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// 成功応答
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("JSON encode error: %v", err)
	}
}
func normalizeString(field *string) *string {
	if field == nil || *field == "" {
		return nil
	}
	return field
}


// ユーザー情報取得
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := usecase.GetUserByID(userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email") // クエリパラメータからemailを取得
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user, err := usecase.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error fetching user by email: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("JSON encode error: %v", err)
	}
}

// ユーザー情報更新
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := usecase.UpdateUser(&user); err != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

