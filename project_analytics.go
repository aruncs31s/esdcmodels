package model

import "time"

type ProjectAnalytics struct {
	ID            uint `gorm:"primaryKey"`
	ProjectID     uint `gorm:"uniqueIndex"`
	TotalViews    int
	TotalLikes    int
	TotalComments int
	AverageRating float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (ProjectAnalytics) TableName() string {
	return "project_analytics"
}