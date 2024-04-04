package entity

import "time"

type Goals struct {
	GoalID      uint      `gorm:"primaryKey;type:serial" json:"goal_id" mapstructure:"goal_id"`
	UserID      string    `gorm:"index" json:"user_id" mapstructure:"user_id"`
	WalletID    string    `gorm:"index" json:"wallet_id" mapstructure:"wallet_id"`
	Amount      float64   `gorm:"not null;default:0" json:"amount" mapstructure:"amount"`
	Description string    `gorm:"size:255" json:"description" mapstructure:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}
