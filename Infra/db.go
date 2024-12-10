package Infra

import (
	"chatSystem/ent"
	"log"

	"github.com/spf13/viper"
)

func getDBClient() (*ent.Client, error) {
	// Viperの設定
	viper.SetConfigName(".env") // ファイル名（拡張子を含まない）
	viper.SetConfigType("env")  // ファイルの形式
	viper.AddConfigPath(".")    // ファイルのパス（カレントディレクトリを指定）

	// .envファイルの読み込み
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUrl := viper.GetString("TIDB_ACC_KEY")
	client, err := ent.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return nil, err
	}
	return client, nil
}
