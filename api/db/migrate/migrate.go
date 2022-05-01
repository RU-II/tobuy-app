package main

import (
	"tobuy-app/api/db"
	"tobuy-app/api/models"

	"gorm.io/gorm"
)

func migrate(dbCon *gorm.DB) {
	// Migration実行
	dbCon.AutoMigrate(&models.User{}, &models.Item{}, &models.Category{}, &models.Group{}, &models.Belong{})
}

func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	// migration実行
	migrate(dbCon)
}
