package repositories

import (
	"gorm.io/gorm"

	"tobuy-app/api/server/models"
)

type UserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetAllUsersByEmail(users *[]models.User, email string) error
	GetUserById(user *models.User, id int) error
	CreateUser(createUser *models.User) error
	UpdateUser(user *models.User, userId int) error
	DeleteUser(userIdf int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

/*
emailに紐づくユーザーを取得
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
userIdに紐づくユーザーを取得
*/
func (ur *userRepository) GetUserById(user *models.User, id int) error {
	if err := ur.db.Where("id=?", id).First(&user).Error; err != nil {
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

/*
ユーザーデータ更新
*/
func (ur *userRepository) UpdateUser(user *models.User, userId int) error {
	if err := ur.db.Where("id=?", userId).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

/*
ユーザーデータ削除
*/
func (ur *userRepository) DeleteUser(userId int) error {
	if err := ur.db.Where("id=?", userId).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
