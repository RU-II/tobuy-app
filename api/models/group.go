package models

type Group struct {
	BaseModel
	Name        string `gorm:"varchar(255)" json:"name,omitempty"`
	Description string `gorm:"type:varchar(255)" json:"description,omitempty"`
	IsPublished bool   `gorm:"type:boolean" json:"is_published,omitempty"`
	LoginId     string `gorm:"varchar(255)" json:"login_id,omitempty"`
	Password    string `gorm:"varchar(255);not null" json:"password,omitempty"`
}
