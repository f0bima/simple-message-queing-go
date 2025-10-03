package dto

type RequestForgotPasswordDto struct {
	Email string `json:"email" binding:"required,email"`
}
