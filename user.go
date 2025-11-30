package model

import (
	"time"
)

type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"column:name;not null" json:"name"`
	Username    string        `gorm:"column:username;unique;not null" json:"username"`
	Email       string        `gorm:"column:email;unique;not null" json:"email"`
	Password    string        `gorm:"column:password;not null" json:"-"`
	Image       *string       `gorm:"column:image" json:"image"`
	Bio         *string       `gorm:"column:bio" json:"bio"`
	Role        string        `gorm:"column:role;not null;default:user" json:"role"`
	Verified    *bool         `gorm:"column:verified;not null;default:false" json:"verified"`
	Status      string        `gorm:"column:status;not null;default:active" json:"status"`
	CreatedAt   time.Time     `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Teams       *[]Teams      `gorm:"many2many:team_members;" json:"teams"`
	Github      *Github       `gorm:"foreignKey:UserID;onDelete:cascade" json:"github"`
	Submissions *[]Submission `gorm:"foreignKey:UserID;onDelete:cascade" json:"submissions"`
	Details     *UserDetails  `gorm:"foreignKey:UserID;references:ID;onDelete:cascade" json:"details"`
	// fixed here:
	Projects      *[]Project     `gorm:"many2many:project_contributors;" json:"projects"`
	LikedProjects *[]Project     `gorm:"many2many:project_likes;" json:"liked_projects"`
	Notifications []Notification `gorm:"foreignKey:UserID;onDelete:cascade" json:"notifications"`
	Location      *Location      `gorm:"foreignKey:UserID;onDelete:cascade" json:"location"`
}

func (User) TableName() string {
	return "users"
}

func (User) GetEssentials() []string {
	return []string{"id",
		"name",
		"username",
		"email",
		"image",
		"role",
	}
}
