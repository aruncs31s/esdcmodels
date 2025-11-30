package model

import (
	"gorm.io/gorm"
)

// Migrate runs auto-migration for all models in the package.
// It creates tables, missing foreign keys, constraints, columns and indexes.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// Core models
		&User{},
		&Project{},
		&Tag{},
		&Technologies{},
		&Comments{},
		&Review{},
		&Order{},
		&Post{},
		&Teams{},
		&Submission{},
		&Notification{},
		&ProjectTemplate{},
		&ProjectStats{},
		&ProjectAnalytics{},

		// Trending models
		&TrendingTech{},
		&TrendingTag{},
		&TrendingCategory{},
		&TrendingAuthor{},
		&TrendingSource{},

		// Integration models
		&ChatBotMessage{},
		&Github{},
		&Ollama{},
	)
}
