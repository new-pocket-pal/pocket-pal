package persistence

import (
	"pocket-pal/src/domain/entity"
	"pocket-pal/src/domain/repository"
	"pocket-pal/src/utils/agents"

	"github.com/sirupsen/logrus"
)

type UserPersistence struct {
	pg  *agents.Postgres
	log *logrus.Logger
}

func (u *UserPersistence) FindUserByID(id uint) (entity.User, error) {
	// implementation of FindUserByID
	// Todo: implement the logic
	return entity.User{}, nil
}

func (u *UserPersistence) RegisterUser(user entity.User) (entity.User, error) {
	// implementation of RegisterUser
	// Todo: implement the logic
	return entity.User{}, nil
}

func (u *UserPersistence) UpdateUser(user entity.User) (entity.User, error) {
	// implementation of UpdateUser
	// Todo: implement the logic
	return entity.User{}, nil
}

func (u *UserPersistence) FindUserByUsername(username string) (entity.User, error) {
	// implementation of FindUserByUsername
	u.log.Info("FindUserByUsername")
	user := new(entity.User)
	err := u.pg.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		u.log.WithError(err).Error("FindUserByUsername")
		return entity.User{}, err
	}

	return *user, nil
}

func (u *UserPersistence) CreateUser(user *entity.User) (entity.User, error) {
	// implementation of CreateUser
	u.log.Info("CreateUser")
	err := u.pg.DB.Create(user).Error
	if err != nil {
		u.log.WithError(err).Error("CreateUser")
		return entity.User{}, err
	}

	return *user, nil
}

func NewUserPersistence(pg *agents.Postgres, log *logrus.Logger) repository.UserRepository {
	return &UserPersistence{
		pg:  pg,
		log: log,
	}
}
