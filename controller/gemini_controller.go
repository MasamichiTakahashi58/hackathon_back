package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/vertexai/genai"
)

const (
	location  = "asia-northeast1"          // デプロイしたリージョン
	modelName = "gemini-1.5-flash-002"     // モデル名
	projectID = "term6-masamichi-takahashi"          // Google CloudプロジェクトID
)

func GenerateContentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Invalid request method:", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// リクエストのパース
	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Gemini APIを呼び出す
	resp, err := generateContent(req.Prompt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating content: %v", err), http.StatusInternalServerError)
		return
	}

	// 結果をJSONで返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func generateContent(prompt string) (interface{}, error) {
	ctx := context.Background()

	// Geminiクライアントの作成
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	// プロンプトを送信してコンテンツ生成
	gemini := client.GenerativeModel(modelName)
	promptObj := genai.Text(prompt)
	resp, err := gemini.GenerateContent(ctx, promptObj)
	if err != nil {
		return nil, fmt.Errorf("error generating content: %w", err)
	}

	return resp, nil
}
