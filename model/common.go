package model

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"not null; index"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
