# ビルドステージ
FROM golang:1.23 as build

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコンテナにコピー
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# アプリケーションをビルド
RUN go build -o main .

# 実行ステージ
FROM gcr.io/distroless/base-debian11

# 作業ディレクトリを設定
WORKDIR /

# ビルド済みのバイナリをコピー
COPY --from=build /app/main .

# ポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
