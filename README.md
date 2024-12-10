# chatSystem

ポートフォリオ向けに作成している
チャットシステムのBackend Rest APIです。
現状の機能は以下の通りです。

- ユーザーの登録(sign up)
- ユーザーのログイン(sign in)


## デプロイ

.envファイルを作成し、以下の通りに記述を行なってください

```.nev
TIDB_ACC_KEY=ent(goのORMライブラリ)形式のDB接続文字列
```
記述後、DBのマイグレーションを行うため、以下のコマンドを実行してください

```bash
go generate ./ent
cd test
go test
```
マイグレーションが完了したら、
コンテナを起動するため、以下のコマンドを実行してください

```bash
docker build -t chat-system .
docker run -d -p 8080:8080 chat-system
```

ローカルで起動する場合は、以下のコマンドを実行してください

```bash
go run main.go
```
