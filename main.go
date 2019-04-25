package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	output, err := GetGroupMessages(42591250, os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}
	GenerateDatapoints(output)
}
