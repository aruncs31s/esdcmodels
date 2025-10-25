package model

import (
	"time"
)

type Project struct {
	ID               uint            `gorm:"primaryKey"`
	Title            string          `gorm:"column:title"`
	Image            *string         `gorm:"column:image"`
	Description      string          `gorm:"column:description"`
	GithubLink       string          `gorm:"column:github_link"`
	LiveURL          *string         `gorm:"column:live_url"`
	Likes            int             `gorm:"column:likes;default:0"`
	Views            int             `gorm:"column:views;default:0"`
	Category         string          `gorm:"column:category;default:'General'"`
	Cost             int             `gorm:"column:cost;default:0"`
	Status           string          `gorm:"column:status;default:'active'"` // active, inactive, archived
	Visibility       int             `gorm:"column:visibility;default:0"`    // 0: public, 1: private
	CreatedBy        uint            `gorm:"column:created_by;not null"`
	ModifiedBy       *uint           `gorm:"column:modified_by"`
	CreatedAt        time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time       `gorm:"column:updated_at;autoUpdateTime"`
	Creator          User            `gorm:"foreignKey:CreatedBy;references:ID" json:"creator"`
	Contributors     *[]User         `gorm:"many2many:project_contributors;" json:"contributors"`
	Tags             *[]Tag          `gorm:"many2many:project_tags;" json:"tags"`
	Technologies     *[]Technologies `gorm:"many2many:project_technologies;" json:"technologies"`
	LikedBy          []User          `gorm:"many2many:project_likes;" json:"liked_by"`
	ViewedBy         []User          `gorm:"many2many:project_views;" json:"viewed_by"`
	ForkedFrom       *uint           `gorm:"column:forked_from"`
	ForkedByProjects *[]Project      `gorm:"foreignKey:ForkedFrom;references:ID;onDelete:cascade"`
	Comments         []Comments      `gorm:"foreignKey:ProjectID;references:ID;onDelete:cascade" json:"comments"`
	Reviews          []Reviews       `gorm:"foreignKey:ProjectID;references:ID;onDelete:cascade" json:"reviews"`
}

func (Project) TableName() string {
	return "projects"
}

func (Project) GetProjectEssentialFields() []string {
	return []string{
		"id",
		"title",
		"image",
		"created_by",
		"status",
		"visibility",
		"likes",
		"views",
		"created_at",
		"updated_at",
	}
}
