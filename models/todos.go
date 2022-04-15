package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uint `gorm:"primary_key;autoincrement" `
	Title     string
	Completed bool
	DueDate   datatypes.Date
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"autoDeleteTime"`
}

type Todos []*Todo
