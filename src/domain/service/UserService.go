package service

import (
	"net/http"
	"pocket-pal/src/domain/command"
	"pocket-pal/src/utils"
	iErr "pocket-pal/src/utils/error"
	"pocket-pal/src/utils/response"

	"github.com/labstack/echo/v4"
)

func (s *Service) RegisterUser(c echo.Context) error {
	// implementation of RegisterUser
	body := new(command.RegisterCommand)

	if err := utils.BindingBody(body, c); err != nil {
		return err
	}

	if body.Username != "" && body.Password != "" {

		existUser, err := s.Facades.UserRepo.FindUserByUsername(body.Username)
		if err == nil && existUser.UserID != "" {
			return c.JSON(http.StatusConflict, response.Response{
				Meta: response.Meta{
					Code: http.StatusConflict,
					Msg:  iErr.ErrUserAlreadyExist.Error(),
					Err:  iErr.ErrUserAlreadyExist,
				},
				Data: nil,
			})
		}

		salt := utils.GenerateSalt()

		hashedPassword := utils.HashPassword(body.Password, salt)

		// generate user entity
		user := command.NewRegisterCommand(body.Username, hashedPassword, body.Email, body.Phone, body.FullName, salt)

		// create user from user entity
		_, err = s.Facades.UserRepo.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Response{
				Meta: response.Meta{
					Code: http.StatusInternalServerError,
					Msg:  iErr.ErrCannotCreateUser.Error(),
					Err:  iErr.ErrCannotCreateUser,
				},
				Data: nil,
			})
		}

		// return success and user data, exclude password and salt, access token
		accessToken, err := utils.GenerateAccessToken(user, s.Facades.Config.JWT.SecretKey, s.Facades.Config.JWT.ExpireTime)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.Response{
				Meta: response.Meta{
					Code: http.StatusInternalServerError,
					Msg:  iErr.ErrCannotGenerateToken.Error(),
					Err:  iErr.ErrCannotGenerateToken,
				},
				Data: nil,
			})
		}

		responseUser := command.RegisterCommandResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			FullName: user.FullName,
		}

		return c.JSON(http.StatusCreated, response.Response{
			Meta: response.Meta{
				Code: http.StatusCreated,
				Msg:  "User created successfully",
				Err:  nil,
			},
			Data: struct {
				User        command.RegisterCommandResponse `json:"user"`
				AccessToken string
			}{
				User:        responseUser,
				AccessToken: accessToken,
			},
		})
	} else {
		return c.JSON(http.StatusBadRequest, response.Response{
			Meta: response.Meta{
				Code: http.StatusBadRequest,
				Msg:  iErr.ErrInvalid.Error(),
				Err:  iErr.ErrInvalid,
			},
		})
	}
}
