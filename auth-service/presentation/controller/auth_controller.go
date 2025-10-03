package auth_controller

import (
	"auth-service/application/usecase"
	interfaces "auth-service/domain/interface"
	"auth-service/presentation/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Producer interfaces.Producer
}

func NewAuthController(p interfaces.Producer) *AuthController {
	return &AuthController{Producer: p}
}

// ForgotPasswordHandler godoc
// @Summary     Forgot Password
// @Description Kirim link reset password ke email
// @Tags        forgot password
// @Accept      json
// @Produce     json
// @Param       request body dto.RequestForgotPasswordDto true "Request"
// @Router       /v1/auth/forgot-password [post]
func (controller *AuthController) ForgotPassword(c *gin.Context) {
	var requestForgotPasswordDto dto.RequestForgotPasswordDto

	if err := c.ShouldBindJSON(&requestForgotPasswordDto); err != nil {
		fmt.Println("Binding Error:", err)

		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body is empty, email required"})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sendForgotPasswordUsecase := usecase.SendForgotPasswordUsecase(controller.Producer)

	if err := sendForgotPasswordUsecase.Execute(requestForgotPasswordDto.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to send email for forgot password link"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requestForgotPasswordDto})
}

// SendOTPHandler godoc
// @Summary     Send OTP
// @Description Kirim OTP ke email
// @Tags        send otp
// @Accept      json
// @Produce     json
// @Param       request body dto.RequestForgotPasswordDto true "Request"
// @Router       /v1/auth/send-otp [post]
func (controller *AuthController) SendOTP(c *gin.Context) {
	var requestEmailOTPDto dto.RequestEmailOTPDto

	if err := c.ShouldBindJSON(&requestEmailOTPDto); err != nil {
		fmt.Println("Binding Error:", err)

		if err.Error() == "EOF" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request body is empty, email required"})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	requestEmailOTPUsecase := usecase.RequestEmailOTPUsecase(controller.Producer)

	if err := requestEmailOTPUsecase.Execute(requestEmailOTPDto.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to send email for forgot password link"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requestEmailOTPDto})
}
