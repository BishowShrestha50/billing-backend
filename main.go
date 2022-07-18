package main

import (
	"os"

	"billing-backend/internal/app/controllers"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error getting env, not coming through %v", err)
	}
	r := controllers.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	err = r.Run(":" + port)
	if err != nil {
		logrus.Error(err)
	}
}
