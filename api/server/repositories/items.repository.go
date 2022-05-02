package repositories

import (
	"tobuy-app/api/server/models"

	"gorm.io/gorm"
)

type ItemsRepository interface {
	GetAllItems(item *[]models.Item, userId int) error
	GetItemById(item *models.Item, id int, userId int) error
}

type itemsRepository struct {
	db *gorm.DB
}

func NewItemsRepository(db *gorm.DB) ItemsRepository {
	return &itemsRepository{db}
}

/*
Itemリストを取得
*/
func (ir *itemsRepository) GetAllItems(items *[]models.Item, userId int) error {
	if err := ir.db.Where("user_id=?", userId).Find(&items).Error; err != nil {
		return err
	}

	return nil
}

func (ir *itemsRepository) GetItemById(item *models.Item, id int, userId int) error {
	if err := ir.db.Where("user_id=?", userId).First(&item, id).Error; err != nil {
		return err
	}

	return nil
}
