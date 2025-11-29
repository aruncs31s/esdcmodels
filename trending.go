package model

type TrendingTech struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	UsageCount int
}

func (TrendingTech) TableName() string {
	return "trending_techs"
}

type TrendingTag struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	UsageCount int
}

func (TrendingTag) TableName() string {
	return "trending_tags"
}

type TrendingCategory struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	UsageCount int
}

func (TrendingCategory) TableName() string {
	return "trending_categories"
}

type TrendingAuthor struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	UsageCount int
}

func (TrendingAuthor) TableName() string {
	return "trending_authors"
}

type TrendingSource struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	UsageCount int
}

func (TrendingSource) TableName() string {
	return "trending_sources"
}
