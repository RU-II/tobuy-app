package models

type User struct {
	BaseModel
	Name     string `gorm:"varchar(255);not null" json:"name,omitempty"`
	Email    string `gorm:"varchar(255);not null;unique" json:"email,omitempty"`
	Password string `gorm:"varchar(255);not null" json:"password,omitempty"`
}
