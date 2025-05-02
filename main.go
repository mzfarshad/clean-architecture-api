package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// TODO: use "github.com/spf13/viper" to handle configurations in different environments
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
}

const (
// userSignUp = "music_store_api/user/auth"
)

func main() {
	//if _, err := models.NewPostgresConnection(); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Successfully connect to postgres")
	//
	//router := gin.Default()
	//router.Use(middleware.LoggerMiddleware())
	//router.Use(middleware.Authenticate())
	//auth := router.Group(userSignUp)
	//
	//auth.POST("/signup", user.SignUp)
	//auth.POST("/signin", user.SignIn)
	//router.Run("localhost:8080")
}
