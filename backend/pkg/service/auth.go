package service

import (
	loginUserDto "github.com/darth-raijin/loop/api/models/dtos/user/login"
	registerUserDto "github.com/darth-raijin/loop/api/models/dtos/user/register"
)

var AuthService authService

type authService struct{}

// Used for registering user
func (authService) CreateUser(user registerUserDto.CreateEventRequest) registerUserDto.CreateUserResponse {

	return registerUserDto.CreateUserResponse{}
}

func (authService) LoginUser(user loginUserDto.LoginUserRequest) loginUserDto.LoginUserResponse {

	return loginUserDto.LoginUserResponse{}
}

func (authService) ResetPassword() {

}

func (authService) FetchUserDetails() {

}

func (authService) verifyUserExists() {

}
