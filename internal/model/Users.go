package model

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DoB       time.Time `json:"dob"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// relations
	UserMediaSocials []UserMediaSocial `json:"user_media_socials" gorm:"foreignKey:UserID;references:ID"`
}

type UserMediaSocial struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// example
type UserSocialMedia struct {
	ID       uint64 `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Title    string `json:"title" gorm:"column:title"`
	Url      string `json:"url" gorm:"column:url"`
}