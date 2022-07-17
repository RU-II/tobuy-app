package logic

import "tobuy-app/api/server/models"

type UsersLogic interface {
	CreateUserResponse(user *models.User) models.BaseUserResponse
}

type usersLogic struct{}

func NewUsersLogic() UsersLogic {
	return &usersLogic{}
}

func (ul *usersLogic) CreateUserResponse(user *models.User) models.BaseUserResponse {
	var userResponse models.BaseUserResponse
	userResponse.BaseModel.ID = user.BaseModel.ID
	userResponse.BaseModel.CreatedAt = user.BaseModel.CreatedAt
	userResponse.BaseModel.UpdatedAt = user.BaseModel.UpdatedAt
	userResponse.BaseModel.DeletedAt = user.BaseModel.DeletedAt
	userResponse.Name = user.Name
	userResponse.Email = user.Email

	return userResponse
}
