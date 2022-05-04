package models

type User struct {
	BaseModel
	Name     string `gorm:"varchar(255);not null" json:"name,omitempty" example:"test user"`
	Email    string `gorm:"varchar(255);not null;unique" json:"email,omitempty" example:"test@example.com"`
	Password string `gorm:"varchar(255);not null" json:"password,omitempty" example:"$2a$10$HgGgk8wX5LjjrssYMuCxQuyeCVooBm7KNlZSNsfemrSpsdKBHDLG."`
}
