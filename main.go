package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	output, err := GetGroupMessages(49006254, os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}
	datapoints := GenerateDatapoints(output)
	payload, _ := json.MarshalIndent(datapoints, "", "     ")

	_ = ioutil.WriteFile("test.json", payload, 0644)
}
