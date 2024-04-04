package entity

import "time"

// GlobalCategory represents the global categories.
type GlobalCategory struct {
	CategoryID  uint      `gorm:"primaryKey"`
	Name        string    `gorm:"unique"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// TableName returns the table name of the global category
func (b *GlobalCategory) TableName() string {
	return "global_categories"
}
