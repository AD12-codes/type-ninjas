package main

import (
	"fmt"

	"github.com/AD12-codes/go-template/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// .env load
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		fmt.Println("Error in loading .env file, using system env vars")
	}
	// server
	server.Run()

}
