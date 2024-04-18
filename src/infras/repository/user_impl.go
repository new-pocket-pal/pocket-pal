package repository

import (
	"pocket-pal/src/domain/entity"
	"pocket-pal/src/domain/repository"
	"pocket-pal/src/facades"
)

type UserImpl struct {
	*facades.Facades
}

func (u *UserImpl) FindUserByID(id uint) (entity.User, error) {
	// implementation of FindUserByID
	// Todo: implement the logic
	return entity.User{}, nil
}

func (u *UserImpl) RegisterUser(user entity.User) (entity.User, error) {
	// implementation of RegisterUser
	// Todo: implement the logic
	return entity.User{}, nil
}

func (u *UserImpl) UpdateUser(user entity.User) (entity.User, error) {
	// implementation of UpdateUser
	// Todo: implement the logic
	return entity.User{}, nil
}

func NewUserImpl(f *facades.Facades) repository.IUser {
	return &UserImpl{
		f,
	}
}
