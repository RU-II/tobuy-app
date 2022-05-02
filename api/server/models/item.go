package models

type Item struct {
	BaseModel
	Name        string `gorm:"type:varchar(63);not null" json:"title,omitempty"`
	Description string `gorm:"type:varchar(255)" json:"description,omitempty"`
	Number      int    `json:"number,omitempty"`
	Counter     string `gorm:"type:varchar(16)" json:"counter,omitempty"`
	Status      int    `gorm:"type:integer;not null" json:"status,omitempty"`
	CategoryId  *int   `gorm:"type:integer;not null;default:0" json:"category_id"`
	// Category    Category `gorm:"foreignKey:CategoryId"`
	UserId *int `gorm:"type:integer;not null;default:0" json:"user_id"`
	// User    User  `gorm:"foreignKey:UserId"`
	GroupId *int `gorm:"type:integer;not null;default:0" json:"group_id"`
	// Group   Group `gorm:"foreignKey:GroupId"`
}

type MutationItemRequest struct {
	Name        string `gorm:"type:varchar(63);not null" json:"name" example:"test item"`
	Description string `gorm:"type:text" json:"description,omitempty" example:"This is a test item"`
	Number      int    `json:"number,omitempty" example:"1"`
	Counter     string `gorm:"type:varchar(16)" json:"counter,omitempty" example:"個"`
	Status      int    `gorm:"type:integer;not null" json:"status,omitempty" example:"0"`
	CategoryId  *int   `gorm:"type:integer;not null;default:0" json:"category_id" example:"1"`
}

type BaseItemResponse struct {
	BaseModel
	Name        string `gorm:"type:varchar(63);not null" json:"name" example:"test item"`
	Description string `gorm:"type:text" json:"description,omitempty" example:"This is a test item"`
	Number      int    `json:"number,omitempty" example:"1"`
	Counter     string `gorm:"type:varchar(16)" json:"counter,omitempty" example:"数"`
	Status      int    `gorm:"type:integer;not null" json:"status,omitempty" example:"0"`
	CategoryId  *int   `gorm:"type:integer;not null;default:0" json:"category_id" example:"1"`
}

type ItemResponse struct {
	Item BaseItemResponse `json:"item"`
}

type AllItemsResponse struct {
	Items []BaseItemResponse `json:"items"`
}
