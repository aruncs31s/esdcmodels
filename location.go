package model

type Location struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	UserID  uint    `gorm:"column:user_id;not null;unique" json:"user_id"`
	City    *string `gorm:"column:city" json:"city"`
	Country *string `gorm:"column:country" json:"country"`
}

func (Location) TableName() string {
	return "locations"
}
