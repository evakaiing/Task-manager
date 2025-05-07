package models

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	IsDone      bool `gorm:"default:false"`
	DueDate     time.Time
	UserID      uint
}
