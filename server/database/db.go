package database

import (
	"fmt"
	"log"
	"os"

	"server/entity"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

const DBMS string = "mysql"

func Connect() *gorm.DB {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("envファイルの読み込みに失敗: ", err)
	}

	CONNECT := getConnectInformation()
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal("DB接続失敗: ", err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Product{})

	fmt.Println("DB接続成功")

	return db
}

func getConnectInformation() string {
	connectInfo := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@" + os.Getenv("DB_PROTOCOL") + "/" + os.Getenv("DB_NAME")
	return connectInfo
}
