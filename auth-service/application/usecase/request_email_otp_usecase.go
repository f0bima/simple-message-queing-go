package usecase

import (
	"auth-service/domain/entity"
	interfaces "auth-service/domain/interface"
	"encoding/json"
	"math/rand"
)

type TRequestEmailOTPUsecase struct {
	Producer interfaces.Producer
}

func RequestEmailOTPUsecase(p interfaces.Producer) *TRequestEmailOTPUsecase {
	return &TRequestEmailOTPUsecase{
		Producer: p,
	}
}

func (uc *TRequestEmailOTPUsecase) Execute(email string) error {
	otpNumber := rand.Intn(9000) + 1000

	otp := entity.Otp{
		Email: email,
		Otp:   otpNumber,
	}

	otpJsonMessage, err := json.Marshal(otp)

	if err != nil {
		return nil
	}

	msg := entity.Message{
		Topic:   "otp",
		Content: string(otpJsonMessage),
	}

	return uc.Producer.Produce(msg)

}
