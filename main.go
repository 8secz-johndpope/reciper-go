package main

import (
	"github.com/SerhiiCho/reciper/src/utils"

	"github.com/SerhiiCho/reciper/src/bootstrap"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError("Error loading .env file", err)
}

func main() {
	bootstrap.IndexPage()
}
