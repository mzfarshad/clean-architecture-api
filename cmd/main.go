package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/mzfarshad/music_store_api/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
}

func main() {
	if _, err := models.NewPostgresConnection(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connect to postgres")
}
