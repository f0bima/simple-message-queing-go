package usecase

import (
	"auth-service/domain/entity"
	interfaces "auth-service/domain/interface"
)

type TSendForgotPasswordUsecase struct {
	Producer interfaces.Producer
}

func SendForgotPasswordUsecase(p interfaces.Producer) *TSendForgotPasswordUsecase {
	return &TSendForgotPasswordUsecase{
		Producer: p,
	}
}

func (uc *TSendForgotPasswordUsecase) Execute(email string) error {
	msg := entity.Message{
		Topic:   "forgot-password",
		Content: email,
	}

	return uc.Producer.Produce(msg)

}
