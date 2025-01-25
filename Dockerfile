# ベースイメージ
FROM --platform=linux/amd64 golang:1.22-alpine AS builder

# 作業ディレクトリの作成
WORKDIR /app

# Go Modulesのキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# ビルド
RUN go build -o server cmd/server/main.go

# エントリーポイント
CMD ["./server"]
