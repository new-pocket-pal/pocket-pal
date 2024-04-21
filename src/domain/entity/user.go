package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID      string `gorm:"primaryKey" json:"user_id" mapstructure:"user_id"`
	Username    string `gorm:"unique;not null;size:50" json:"username" mapstructure:"username"`
	Password    string `gorm:"not null;size:255" json:"password" mapstructure:"password"`
	Salt        string
	Email       string    `gorm:"unique;not null;size:100" json:"email" mapstructure:"email"`
	Phone       string    `gorm:"unique;not null;size:15" json:"phone" mapstructure:"phone"`
	FullName    string    `gorm:"size:100" json:"full_name" mapstructure:"full_name"`
	DateOfBirth time.Time `gorm:"type:date" json:"date_of_birth" mapstructure:"date_of_birth"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" mapstructure:"updated_at"`
}

// TableName is a function to get the table name
func (User) TableName() string {
	return "users"
}

// BeforeCreate is a function to set the created_at field
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var userId string
	tx.Raw("SELECT generate_user_id();").Scan(&userId)
	u.UserID = userId
	u.CreatedAt = time.Now()
	// set user_id
	return nil
}

// BeforeUpdate is a function to set the updated_at field
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
