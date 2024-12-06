package models

import "time"

type Article struct {
	ArticleID string     `gorm:"primaryKey" json:"article_id"`
	CategoryID string    `gorm:"foreignKey" json:"category_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   string    `gorm:"foreignKey" json:"author_id"`
	ViewCount  int       `json:"view_count"`
	Slug       string    `json:"slug"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}