package model

type Corse struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Views       uint   `gorm:"default:0" json:"views"`
	Likes       uint   `gorm:"default:0" json:"likes"`
	Enrollments uint   `gorm:"default:0" json:"enrollments"`
}

func (Corse) TableName() string {
	return "courses"
}
