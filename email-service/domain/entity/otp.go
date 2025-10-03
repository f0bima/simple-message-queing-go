package entity

type Otp struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}
