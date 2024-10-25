package models

import "time"

type Newsletter struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"default:'Newsletter'" json:"title"`
	Attachment string    `json:"attachment"`
	Html       string    `gorm:"type:text" json:"html"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	ScheduledDate *time.Time `gorm:"type:date" json:"scheduledDate"`
	Scheduled     bool       `gorm:"default:false" json:"scheduled"`

	CategoryID *uint     `gorm:"index" json:"categoryId,omitempty"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`

	Recipients []Recipient `gorm:"many2many:newsletter_recipients;" json:"recipients,omitempty"`
}

type Category struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string       `gorm:"unique" json:"title"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updatedAt"`
	Newsletters []Newsletter `gorm:"foreignKey:CategoryID" json:"newsletters,omitempty"`
}

type Recipient struct {
	ID            string       `gorm:"primaryKey" json:"id"`
	Email         string       `gorm:"unique" json:"email"`
	UnsuscribeUrl string       `gorm:"type:varchar(255)" json:"unsuscribeUrl"`
	CreatedAt     time.Time    `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime" json:"updatedAt"`
	Newsletters   []Newsletter `gorm:"many2many:newsletter_recipients;" json:"newsletters,omitempty"`
}
