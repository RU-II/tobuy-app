package models

type User struct {
	BaseModel
	Name     string `gorm:"varchar(255);not null" json:"name,omitempty" example:"test user"`
	Email    string `gorm:"varchar(255);not null;unique" json:"email,omitempty" example:"test@example.com"`
	Password string `gorm:"varchar(255);not null" json:"password,omitempty" example:"$2a$10$HgGgk8wX5LjjrssYMuCxQuyeCVooBm7KNlZSNsfemrSpsdKBHDLG."`
}

type BaseUserResponse struct {
	BaseModel
	Name  string `gorm:"size:255" json:"name,omitempty" example:"test user"`
	Email string `gorm:"size:255;not null;unique" json:"email,omitempty" example:"test@example.com"`
}

type UserResponse struct {
	User BaseUserResponse `json:"user"`
}

type UpdateUserRequest struct {
	Name  string `gorm:"varchar(255);not null" json:"name,omitempty" example:"test user"`
	Email string `gorm:"varchar(255);not null;unique" json:"email,omitempty" example:"test@example.com"`
}

type UpdatePasswordRequest struct {
	Password    string `gorm:"varchar(255);not null" json:"password,omitempty" example:"password"`
	NewPassword string `gorm:"varchar(255);not null" json:"new_password,omitempty" example:"newPassword"`
}

type DeleteUserRequest struct {
	Password string `gorm:"varchar(255);not null" json:"password,omitempty" example:"password"`
}
