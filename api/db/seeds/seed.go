package main

import (
	"fmt"
	"math/rand"
	"tobuy-app/api/db"
	"tobuy-app/api/models"

	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// user 10人, group 5個 作成

func userSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		user := models.User{
			Name:     "ユーザー" + strconv.Itoa(i+1),
			Email:    "sample" + strconv.Itoa(i+1) + "@gmail.com",
			Password: string(hash),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func groupSeeds(db *gorm.DB) error {
	for i := 0; i < 5; i++ {
		group := models.Group{
			Name:        "グループ" + strconv.Itoa(i+1),
			Description: "sample" + strconv.Itoa(i+1),
		}

		if err := db.Create(&group).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func belongSeeds(db *gorm.DB) error {
	for i := 0; i < 20; i++ {
		userId := (i % 10) + 1
		groupId := rand.Intn(5-1) + 1
		belong := models.Belong{
			UserId:  &userId,
			GroupId: &groupId,
		}

		if err := db.Create(&belong).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func itemSeeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			categoryId := rand.Intn(5) + 1
			status := rand.Intn(2)
			userId := i + 1
			userItem := models.Item{
				Name:        "ユーザー" + strconv.Itoa(i+1) + "アイテム",
				Description: "アイテム" + strconv.Itoa(j+1),
				Number:      rand.Intn(8),
				Counter:     "個",
				Status:      status,
				CategoryId:  &categoryId,
				UserId:      &userId,
			}
			if err := db.Create(&userItem).Error; err != nil {
				fmt.Printf("%+v", err)
			}
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			categoryId := rand.Intn(5) + 1
			status := rand.Intn(2)
			groupId := i + 1
			groupItem := models.Item{
				Name:        "グループ" + strconv.Itoa(i+1) + "アイテム",
				Description: "アイテム" + strconv.Itoa(j+1),
				Number:      rand.Intn(8),
				Counter:     "個",
				Status:      status,
				CategoryId:  &categoryId,
				GroupId:     &groupId,
			}
			if err := db.Create(&groupItem).Error; err != nil {
				fmt.Printf("%+v", err)
			}
		}
	}
	return nil
}

func categorySeeds(db *gorm.DB) error {
	categories := []string{"food", "commodity", "clothes", "hobby", "others"}
	for _, category := range categories {

		if err := db.Create(&models.Category{Name: category}).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func main() {
	dbCon := db.Init()
	// dBを閉じる
	defer db.CloseDB(dbCon)

	if err := userSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
		return
	}

	if err := groupSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
		return
	}

	if err := belongSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
		return
	}

	if err := itemSeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
		return
	}

	if err := categorySeeds(dbCon); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
