package models

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(63);not null" json:"name"`
}
