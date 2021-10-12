package main

import (
	"fmt"
	"server/database"
	"server/entity"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func seeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		todo := entity.Todo{Name: "課題", Memo: "メモ", Status: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func main() {
	db := database.Connect()
	defer db.Close()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
