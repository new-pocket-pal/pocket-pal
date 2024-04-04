package entity

import "time"

// Budgets represents the budgets.
type Budgets struct {
	BudgetID    uint      `gorm:"primaryKey;type:serial" json:"budget_id" mapstructure:"budget_id"`
	UserID      uint      `gorm:"index" json:"user_id" mapstructure:"user_id"`
	WalletID    uint      `gorm:"index" json:"wallet_id" mapstructure:"wallet_id"`
	Amount      float64   `gorm:"not null;default:0" json:"amount" mapstructure:"amount"`
	StartDate   time.Time `gorm:"type:date" json:"start_date" mapstructure:"start_date"`
	EndDate     time.Time `gorm:"type:date" json:"end_date" mapstructure:"end_date"`
	Description string    `gorm:"size:255" json:"description" mapstructure:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}

// Budgets table name of the database.
func (Budgets) TableName() string {
	return "budgets"
}
