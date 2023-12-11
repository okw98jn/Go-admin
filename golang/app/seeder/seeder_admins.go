package main

import (
	"app/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func seedAdmins(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE admins")
	// レコード数
	dataLength := 20000
	admins := make([]model.Admin, dataLength)
	for i := 0; i < dataLength; i++ {
		admins[i] = model.Admin{
			Name:      fmt.Sprintf("Admin%d", i+1),
			LoginId:   fmt.Sprintf("login%d", i+1),
			Rore:      i%2 == 0,
			Status:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Now(),
		}
	}

	// 処理時間短縮のためにトランザクションを使用
	tx := db.Begin()
	for _, admin := range admins {
		tx.Create(&admin)
	}
	tx.Commit()
	fmt.Println("admin seeder finished")
}
