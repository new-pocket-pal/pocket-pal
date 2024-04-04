package entity

import "time"

// Wallet struct is a representation of the wallet table in the database
type Wallet struct {
	WalletID   string    `gorm:"primaryKey" json:"wallet_id" mapstructure:"wallet_id"`
	UserID     string    `gorm:"index" json:"user_id" mapstructure:"user_id"`
	Name       string    `gorm:"not null;size:100" json:"name" mapstructure:"name"`
	Balance    float64   `gorm:"not null;default:0" json:"balance" mapstructure:"balance"`
	Currency   string    `gorm:"size:3" json:"currency" mapstructure:"currency"`
	IsMain     bool      `gorm:"not null;default:false" json:"is_main" mapstructure:"is_main"`
	WalletType string    `gorm:"not null;size:30" json:"wallet_type" mapstructure:"wallet_type"`
	IsArchived bool      `gorm:"not null;default:false" json:"is_archived" mapstructure:"is_archived"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}

// TableName returns the table name of the wallet
func (Wallet) TableName() string {
	return "wallets"
}
