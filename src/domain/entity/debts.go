package entity

import "time"

// Debts represents the debts between users.
type Debts struct {
	DebtID         uint      `gorm:"primaryKey;type:serial" json:"debt_id" mapstructure:"debt_id"`
	DebtorUserID   string    `gorm:"index" json:"debtor_user_id" mapstructure:"debtor_user_id"`
	CreditorUserID string    `gorm:"index" json:"creditor_user_id" mapstructure:"creditor_user_id"`
	Amount         float64   `gorm:"not null;default:0" json:"amount" mapstructure:"amount"`
	Description    string    `gorm:"size:255" json:"description" mapstructure:"description"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}

// Debts table name of the database.
func (Debts) TableName() string {
	return "debts"
}
