package error

import (
	"errors"
)

var (
	ErrUserAlreadyExist    = errors.New("user already existed")
	ErrCannotCreateUser    = errors.New("cannot create user")
	ErrCannotGenerateToken = errors.New("cannot generate token")
)
