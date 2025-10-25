package model

import "time"

type Reviews struct {
	ID        uint      `gorm:"primaryKey"`
	ProjectID uint      `gorm:"column:project_id;not null"`
	UserID    uint      `gorm:"column:user_id;not null"`
	Rating    int       `gorm:"column:rating;not null"`
	Comment   string    `gorm:"column:comment"`
	User      *User     `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"user"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Reviews) TableName() string {
	return "reviews"
}
