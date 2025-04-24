package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mzfarshad/music_store_api/internal/api/handler/user"
	"github.com/mzfarshad/music_store_api/internal/api/middleware"
	"github.com/mzfarshad/music_store_api/internal/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
}

const (
	userSignUp = "music_store_api/user/auth"
)

func main() {
	if _, err := models.NewPostgresConnection(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connect to postgres")

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	auth := router.Group(userSignUp)

	auth.POST("/signup", user.SignUp)
	router.Run("localhost:8080")
}
