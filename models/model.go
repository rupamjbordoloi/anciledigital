package models

import (
	"time"

	"gorm.io/gorm"
)

type IDColumn struct {
	ID uint `gorm:"primarykey" json:"id"`
}

type DateColumn struct {
	CreatedAt time.Time      `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggertype:"string"`
}
