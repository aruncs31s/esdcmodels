package model

import "time"

type ProjectTemplate struct {
	ID          uint `gorm:"primaryKey"`
	ProjectID   uint `gorm:"index"`
	Name        string
	Description string
	CreatedBy   uint `gorm:"index"`
	IsPublic    bool `gorm:"default:false"`
	UsageCount  int  `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// Relationships
	Creator      *User
	Technologies *[]Technologies `gorm:"many2many:template_technologies;"`
	Tags         *[]Tag          `gorm:"many2many:template_tags;"`
}

func (ProjectTemplate) TableName() string {
	return "project_templates"
}
