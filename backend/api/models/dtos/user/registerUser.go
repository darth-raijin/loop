package dtos

type RegisterUserDto struct {
	Name     string `json:"name" validate:"required" example:"Go Crash Course"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
