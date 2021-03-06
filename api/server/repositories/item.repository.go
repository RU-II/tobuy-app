package repositories

import (
	"tobuy-app/api/server/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems(item *[]models.Item, userId int) error
	GetItemById(item *models.Item, id int, userId int) error
	GetLastItem(item *models.Item, userId int) error
	CreateItem(item *models.Item) error
	DeleteItem(id int, userId int) error
	UpdateItem(item *models.Item, id int, userId int) error
}

type itemsRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
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

func (ir *itemsRepository) GetLastItem(item *models.Item, userId int) error {
	if err := ir.db.Where("user_id=?", userId).Last(&item).Error; err != nil {
		return err
	}

	return nil
}

func (ir *itemsRepository) CreateItem(item *models.Item) error {
	if err := ir.db.Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func (ir *itemsRepository) DeleteItem(id int, userId int) error {
	if err := ir.db.Where("id=? AND user_id=?", id, userId).Delete(&models.Item{}).Error; err != nil {
		return err
	}

	// deleteすると必ず、RowsAffected < 1になるのでコメントアウト
	// https://stackoverflow.com/questions/67154864/how-to-handle-gorm-error-at-delete-function
	// if ir.db.Error != nil {
	// 	return ir.db.Error
	// } else {
	// 	if ir.db.RowsAffected < 1 {
	// 		// return nil
	// 		return errors.Errorf("id=%w のアイテムデータが存在しません。", id)
	// 	}
	// }

	return nil
}

func (ir *itemsRepository) UpdateItem(item *models.Item, id int, userId int) error {
	if err := ir.db.Where("id=? AND user_id=?", id, userId).Updates(&item).Error; err != nil {
		return err
	}
	return nil
}
