package initialize

import (
	"fmt"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found or error loading .env file")
	}
}
