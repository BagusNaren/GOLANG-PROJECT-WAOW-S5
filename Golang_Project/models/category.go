package models

import "time"

type Category struct {
	CategoryID string     `gorm:"primaryKey" json:"category_id"`
	Name       string     `json:"name"`
	Description string    `json:"description"`
	Slug       string     `json:"slug"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}