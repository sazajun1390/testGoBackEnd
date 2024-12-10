package todo

import (
	"chatSystem/ent"
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func TestCreateTable(t *testing.T) {
	// インメモリーのSQLiteデータベースを持つent.Clientを作成します。
	//client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	viper.SetConfigName(".env") // ファイル名（拡張子を含まない）
	viper.SetConfigType("env")  // ファイルの形式
	viper.AddConfigPath(".")    // ファイルのパス（カレントディレクトリを指定）

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUrl := viper.GetString("TIDB_ACC_KEY")
	client, err := ent.Open("mysql", dbUrl)
	if err != nil {
		t.Fatalf("failed opening connection to sqlite: %v", err)
		//log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// 自動マイグレーションツールを実行して、すべてのスキーマリソースを作成します。
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
		//log.Fatalf("failed creating schema resources: %v", err)
	}
	// 出力します。
}
