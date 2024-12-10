# ベースイメージ
FROM golang:1.23.3-alpine

# 作業ディレクトリの設定
WORKDIR /app

# Goモジュールを有効化
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# アプリケーションソースコードをコピー
COPY . ./

# アプリケーションのビルド
RUN go build -o main .

# コンテナ起動時に実行するコマンド
CMD ["./main"]

# コンテナの公開ポート
EXPOSE 8080
