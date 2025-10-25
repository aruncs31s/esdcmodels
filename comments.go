package model

import "time"

type Comments struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"column:project_id;not null"`
	UserID    uint      `gorm:"column:user_id;not null"`
	Content   string    `gorm:"column:content;not null"`
	User      *User     `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"user"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Comments) TableName() string {
	return "comments"
}
