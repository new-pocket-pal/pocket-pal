package repository

import "pocket-pal/src/domain/entity"

type UserRepository interface {
	FindUserByUsername(username string) (entity.User, error)
	FindUserByID(userId uint) (entity.User, error)
	RegisterUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	CreateUser(user *entity.User) (entity.User, error)
}
