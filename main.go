package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	output, err := GetGroupMessages(32945761, os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}
	fmt.Println(len(output))
	datapoints := GenerateDatapoints(output)
	var newDatapoints []int
	for i := 0; i < len(datapoints); i++ {
		if i == 0 {
			newDatapoints = append(newDatapoints, datapoints[i])
		} else {
			newDatapoints = append(newDatapoints, datapoints[i]+newDatapoints[i-1])
		}
	}

	//payload, _ := json.MarshalIndent(newDatapoints, "", "     ")

	var s string
	for i := 0; i < len(newDatapoints); i++ {
		s = s + fmt.Sprintf("%d\n", newDatapoints[i])
	}
	_ = ioutil.WriteFile("test.json", []byte(s), 0644)
}
