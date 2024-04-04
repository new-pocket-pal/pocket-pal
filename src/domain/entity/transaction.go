package entity

import "time"

// Transactions struct is a representation of the transactions table in the database
type Transactions struct {
	TransactionID   string    `gorm:"primaryKey" json:"transaction_id" mapstructure:"transaction_id"`
	UserID          string    `gorm:"index" json:"user_id" mapstructure:"user_id"`
	WalletID        string    `gorm:"index" json:"wallet_id" mapstructure:"wallet_id"`
	CategoryID      uint      `gorm:"index" json:"category_id" mapstructure:"category_id"`
	Amount          float64   `gorm:"not null;default:0" json:"amount" mapstructure:"amount"`
	TransactionType string    `gorm:"not null;size:30" json:"transaction_type" mapstructure:"transaction_type"`
	Description     string    `gorm:"size:255" json:"description" mapstructure:"description"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}

// TableName returns the table name of the transactions
func (Transactions) TableName() string {
	return "transactions"
}
