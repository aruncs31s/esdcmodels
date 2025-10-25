package model

type Github struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	UserID   uint   `gorm:"column:user_id"`
	Username string `gorm:"column:username;unique;not null"`
}

func (Github) TableName() string {
	return "github"
}
