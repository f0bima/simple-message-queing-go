package main

import (
	_ "auth-service/docs"
	kafka_producer "auth-service/infrastructure/kafka"
	auth_controller "auth-service/presentation/controller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          Auth Service
// @version        1.0
// @description    Dokumentasi api tentang penggunaan kafka producer dengan studi kasus auth service(Forgot password dan Send OTP)
// @contact.name   Fauzan Bima
// @contact.url    https://www.f0bima.com
// @contact.email  fbimapk@gmail.io

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// broker := os.Getenv("BROKER_URL")

	config := kafka_producer.KafkaProducerConfig{
		Broker:           os.Getenv("BROKER_URL"),
		SecurityProtocol: os.Getenv("SECURITY_PROTOCOL"),
		SaslMechanism:    os.Getenv("SASL_MECHANISM"),
		Username:         os.Getenv("SASL_USERNAME"),
		Password:         os.Getenv("SASL_PASSWORD"),
	}
	producer, err := kafka_producer.NewKafkaProducer(config)

	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s\n", err)
		return
	}
	defer producer.Close()

	r := gin.Default()

	authController := auth_controller.NewAuthController(producer)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/v1/auth/forgot-password", authController.ForgotPassword)
	r.POST("/v1/auth/send-otp", authController.SendOTP)

	r.Run()
}
