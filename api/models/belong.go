package models

type Belong struct {
	BaseModel
	UserId *int `gorm:"type:integer;not null;default:0" json:"user_id"`
	// User    User  `gorm:"foreignKey:UserId"`
	GroupId *int `gorm:"type:integer;not null;default:0" json:"group_id"`
	// Group   Group `gorm:"foreignKey:GroupId"`
}
