package models

import (
	"time"
)

// gorm.Modelで置換
type BaseModel struct {
	ID        int        `gorm:"type:serial;primaryKey" json:"id" example:"1"`
	CreatedAt *time.Time `json:"created_at" example:"2022-05-01T17:23:17.494039+09:00"`
	UpdatedAt *time.Time `json:"updated_at" example:"2022-05-01T17:23:17.494039+09:00"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:""`
}
