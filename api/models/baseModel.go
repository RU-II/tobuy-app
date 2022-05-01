package models

import (
	"time"
)

// gorm.Modelで置換
type BaseModel struct {
	ID        int        `gorm:"type:serial;primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
