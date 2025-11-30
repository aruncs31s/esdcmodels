package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Review represents a project review in the system
type Review struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	ProjectID uint       `gorm:"index" json:"project_id"`
	UserID    uint       `gorm:"index" json:"user_id"`
	Rating    float64    `gorm:"column:rating" json:"rating" validate:"required,min=1,max=5"`
	Comment   string     `gorm:"type:text" json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"` // Soft delete

	// Relations
	Project *Project `gorm:"foreignKey:ProjectID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	User    *User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName specifies the table name for the Review model
func (Review) TableName() string {
	return "project_reviews"
}

// BeforeCreate hook for Review
func (r *Review) BeforeCreate(tx *gorm.DB) error {
	if r.Rating < 1 || r.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return nil
}
