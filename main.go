package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	output, err := GetGroupMessages(44330425, os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}
	x := SeperateUsers(output)
	y := make(map[string][]int)
	for person := range x {
		for i := 0; i < len(x[person]); i++ {
			if i == 0 {
				y[person] = append(y[person], x[person][i])
			} else {
				y[person] = append(y[person], x[person][i]+y[person][i-1])
			}
		}
	}

	maxLength := 0
	for person := range y {
		length := len(y[person])
		fmt.Println(length)
		if length > maxLength {
			maxLength = length
		}
	}

	var s string
	for person := range y {
		s = s + person + ","
	}
	s += "\n"
	for i := 0; i < maxLength; i++ {
		for person := range y {
			if i < len(y[person]) {
				additiveString := fmt.Sprintf("%d,", y[person][i])
				s = s + additiveString
				if y[person][i] == 0 && i > 200 {
					fmt.Println(y[person][i])
					panic(":)")
				}
			} else {
				s = s + fmt.Sprintf("%d,", y[person][len(y[person])-1])
			}
		}
		s += "\n"
	}

	//payload, _ := json.MarshalIndent(y, "", "    ")
	_ = ioutil.WriteFile("data", []byte(s), 0644)
}
