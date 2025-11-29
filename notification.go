package model

import "time"

type Notification struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index"`
	Type        string // like, comment, follow, milestone
	Title       string
	Message     string
	ProjectID   uint `gorm:"index"`
	TriggeredBy uint
	IsRead      bool `gorm:"default:false"`
	CreatedAt   time.Time
	// Relationships
	User            *User
	Project         *Project
	TriggeredByUser *User `gorm:"foreignKey:TriggeredBy"`
}

func (Notification) TableName() string {
	return "notifications"
}