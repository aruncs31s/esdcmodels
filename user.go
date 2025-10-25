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

type UserDetails struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	UserID        uint    `gorm:"column:user_id;not null;unique" json:"user_id"`
	Twitter       *string `gorm:"column:twitter" json:"twitter"`
	Location      *string `gorm:"column:location" json:"location"`
	LinkedIn      *string `gorm:"column:linkedin" json:"linkedin"`
	Facebook      *string `gorm:"column:facebook" json:"facebook"`
	Instagram     *string `gorm:"column:instagram" json:"instagram"`
	StackOverflow *string `gorm:"column:stackoverflow" json:"stackoverflow"`
	Website       *string `gorm:"column:website" json:"website"`
}

func (UserDetails) TableName() string {
	return "user_details"
}

type Location struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	UserID  uint    `gorm:"column:user_id;not null;unique" json:"user_id"`
	City    *string `gorm:"column:city" json:"city"`
	Country *string `gorm:"column:country" json:"country"`
}

func (Location) TableName() string {
	return "locations"
}

func (User) TableName() string {
	return "users"
}
