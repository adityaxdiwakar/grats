package main

import (
	"fmt"
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
	seperateOutput := SeperateUsers(output)
	seperatedMessages := make(map[string][]int)
	for person := range seperateOutput {
		for i := 0; i < len(seperateOutput[person]); i++ {
			if i == 0 {
				seperatedMessages[person] = append(seperatedMessages[person], seperateOutput[person][i])
			} else {
				seperatedMessages[person] = append(seperatedMessages[person], seperateOutput[person][i]+seperatedMessages[person][i-1])
			}
		}
	}

	maxLength := 0
	for person := range seperatedMessages {
		length := len(seperatedMessages[person])
		if length > maxLength {
			maxLength = length
		}
	}

	var namelessArray [][]int
	var people []string
	for person := range seperatedMessages {
		people = append(people, person)
		namelessArray = append(namelessArray, seperatedMessages[person])
	}
	var s string
	for i := 0; i < len(people); i++ {
		s += people[i] + ","
	}
	s += "\n"
	for i := 0; i < maxLength; i++ {
		fmt.Println(100 * i / maxLength)
		for h := 0; h < len(namelessArray); h++ {
			if i < len(namelessArray[h]) {
				s += fmt.Sprintf("%d,", namelessArray[h][i])
			} else {
				s += fmt.Sprintf("%d,", namelessArray[h][len(namelessArray[h])-1])
			}
		}
		s += "\n"
	}

	//payload, _ := json.MarshalIndent(y, "", "    ")
	_ = ioutil.WriteFile("data.csv", []byte(s), 0644)
}
