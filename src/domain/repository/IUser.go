package repository

import "pocket-pal/src/domain/entity"

type IUser interface {
	FindUserByID(userId uint) (entity.User, error)
	RegisterUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
}
