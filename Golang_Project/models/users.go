package models

import "time"

type User struct {
	UserID           string    `gorm:"primaryKey" json:"user_id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"-"`
	DisplayName      string    `json:"display_name"`
	Bio              string    `json:"bio"`
	ProfilePictureURL string   `json:"profile_picture_url"`
	RegistrationDate time.Time `json:"registration_date"`
	Role             string    `json:"role"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}