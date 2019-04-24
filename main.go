package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	output, err := getGroupInformation(32945761)
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(output, "", "    ")
	_ = ioutil.WriteFile("test.json", data, 0777)
}
