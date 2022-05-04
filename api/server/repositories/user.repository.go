package repositories

import (
	"gorm.io/gorm"

	"tobuy-app/api/server/models"
)

type UserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetAllUsersByEmail(users *[]models.User, email string) error
	CreateUser(createUser *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

/*
emailに紐づくユーザーリストを取得
*/
func (ur *userRepository) GetUserByEmail(user *models.User, email string) error {
	// db := db.GetDB()
	if err := ur.db.Where("email=?", email).First(&user).Error; err != nil {
		return err
	}

	return nil
}

/*
emailに紐づくユーザーリストを取得
*/
func (ur *userRepository) GetAllUsersByEmail(users *[]models.User, email string) error {
	// db := db.GetDB()
	if err := ur.db.Where("email=?", email).Find(&users).Error; err != nil {
		return err
	}

	return nil
}

/*
ユーザーデータ新規登録
*/
func (ur *userRepository) CreateUser(createUser *models.User) error {
	// db := db.GetDB()
	if err := ur.db.Create(&createUser).Error; err != nil {
		return err
	}

	return nil
}