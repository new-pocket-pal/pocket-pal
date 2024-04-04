package entity

import "time"

// Category struct is a representation of the category table in the database
type CustomCategory struct {
	CategoryID uint      `gorm:"primaryKey;type:serial" json:"category_id" mapstructure:"category_id"`
	Name       string    `gorm:"not null;size:100" json:"name" mapstructure:"name"`
	Type       string    `gorm:"not null;size:50" json:"type" mapstructure:"type"`
	UserID     string    `gorm:"index" json:"user_id" mapstructure:"user_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}

// TableName returns the table name of the category
func (b *CustomCategory) TableName() string {
	return "categories"
}
