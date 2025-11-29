package model

import "time"

type ProjectStats struct {
	ID            uint    `gorm:"primaryKey"`
	ProjectID     uint    `gorm:"uniqueIndex"`
	ViewCount     int     `gorm:"default:0"`
	LikeCount     int     `gorm:"default:0"`
	CommentCount  int     `gorm:"default:0"`
	ReviewCount   int     `gorm:"default:0"`
	AverageRating float64 `gorm:"default:0"`
	UpdatedAt     time.Time
}

func (ProjectStats) TableName() string {
	return "project_stats"
}