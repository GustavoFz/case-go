package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func InitEnv() map[string]string {

	envs, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error loading .env")
	}

	fmt.Println(".env successfully loaded")
	return envs
}
