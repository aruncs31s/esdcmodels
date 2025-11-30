package model

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
