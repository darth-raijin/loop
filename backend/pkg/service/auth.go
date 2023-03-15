package service

import (
	"time"

	errorDto "github.com/darth-raijin/loop/api/models/dtos/error"
	loginUserDto "github.com/darth-raijin/loop/api/models/dtos/user/login"
	registerUserDto "github.com/darth-raijin/loop/api/models/dtos/user/register"
	"github.com/darth-raijin/loop/api/models/entities"
	"github.com/darth-raijin/loop/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

var AuthService authService

type authService struct{}

// Used for registering user
func (authService) CreateUser(user registerUserDto.CreateEventRequest) (registerUserDto.CreateUserResponse, errorDto.DomainErrorWrapper) {
	userEntity := entities.User{
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return registerUserDto.CreateUserResponse{}, errorDto.DomainErrorWrapper{
			Timestamp: time.Now(),
			Errors: []errorDto.DomainError{
				{
					Message: err.Error(),
				},
			},
		}
	}

	userEntity.Password = string(hashed)

	// Check if EMAIL and USERNAME is distinct
	result := repository.GormDB.Create(&userEntity)
	if result.Error != nil {
		// Handle database error

		return registerUserDto.CreateUserResponse{}, errorDto.DomainErrorWrapper{
			Timestamp: time.Now(),
			Errors: []errorDto.DomainError{
				{
					Message: result.Error.Error(),
				},
			},
		}
	}

	return registerUserDto.CreateUserResponse{}, errorDto.DomainErrorWrapper{}
}

func (authService) LoginUser(user loginUserDto.LoginUserRequest) loginUserDto.LoginUserResponse {

	return loginUserDto.LoginUserResponse{}
}

func (authService) ResetPassword() {

}

func (authService) FetchUserDetails() {

}
