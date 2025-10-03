package dto

type RequestEmailOTPDto struct {
	Email string `json:"email" binding:"required,email"`
}
