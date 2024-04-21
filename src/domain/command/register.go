package command

import "pocket-pal/src/domain/entity"

type RegisterCommand struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
}

type RegisterCommandResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
}

// NewRegisterCommand to UserEntity
func NewRegisterCommand(username, password, email, phone, fullName, salt string) *entity.User {
	return &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
		FullName: fullName,
		Salt:     salt,
	}
}
