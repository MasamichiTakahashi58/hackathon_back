package usecase

import (
    "hackathon_back/dao"
    "hackathon_back/model"

    "os"
	"path/filepath"
	"time"
    "fmt"
)

// ユーザーの作成
func CreateUser(user *model.User) error {
	return dao.CreateUser(user)
}

// ユーザーの取得
func GetUserByID(userID int) (*model.User, error) {
    return dao.GetUserByID(userID)
}
func GetUserByEmail(email string) (*model.User, error) {
	return dao.GetUserByEmail(email)
}

// ユーザーの更新
func UpdateUser(user *model.User) error {
    return dao.UpdateUser(user)
}


func SaveUserImage(userID int, imageType, originalFilename string, fileData []byte) (string, error) {
	// 保存先ディレクトリを決定
	saveDir := filepath.Join("uploads", imageType)
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// タイムスタンプを使ってファイル名を一意に生成
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("user%d-%s-%d%s", userID, imageType, timestamp, filepath.Ext(originalFilename))
	filePath := filepath.Join(saveDir, filename)

	// ファイルを書き込む
	if err := os.WriteFile(filePath, fileData, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// データベースにパスを保存
	dbFilePath := fmt.Sprintf("/uploads/%s/%s", imageType, filename)
	column := "profile_image"
	if imageType == "header" {
		column = "header_image"
	}

	if err := dao.UpdateUserImage(userID, column, dbFilePath); err != nil {
		return "", fmt.Errorf("failed to update database: %w", err)
	}

	return dbFilePath, nil
}
