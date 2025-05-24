FROM golang:1.24

WORKDIR /app

# ホットリロード用のツールとしてairをインストール
RUN go install github.com/air-verse/air@latest
RUN go install golang.org/x/lint/golint@latest

COPY go/. ./
RUN go mod download