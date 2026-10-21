package main

import (
	"fmt"

	"Testify-B/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../../.env")
	config.LoadConfig()
	fmt.Println("Goodbye World")
}
