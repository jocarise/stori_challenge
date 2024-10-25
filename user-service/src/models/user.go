package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Role      string    `gorm:"default:'ADMIN'" json:"-"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}
