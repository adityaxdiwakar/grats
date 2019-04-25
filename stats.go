package main

import (
	"fmt"
	"math"
)

func GenerateDatapoints(messages []Message) {
	fmt.Println(len(messages))
	var counts []int
	counter := 0
	last_time := float64(-1)
	for i := 0; i < len(messages); i++ {
		time := messages[i].CreatedAt
		hour := math.Floor(float64(time / (3600)))
		if hour != last_time {
			counts = append(counts, counter)
			counter = 0
			last_time = hour
		} else {
			counter = counter + 1
		}
	}
	counts = append(counts, counter)
	fmt.Println(counts)
}
