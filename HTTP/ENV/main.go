/*
	in .env file:
	OPENWEATHERMAP_API_KEY=abc123xyz456

	package:
	github.com/joho/godotenv
	os

*/

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENWEATHERMAP_API_KEY is not set in the .env file")
		return
	}
	fmt.Println(apiKey)
}
