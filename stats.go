package main

import (
	"math"
)

func GenerateDatapoints(messages []Message, startTime int) []int {
	var timeCorrelationPoints []int
	timeGrouping := 3600
	counter := 0
	timeHolding := math.Floor(float64(startTime / timeGrouping))
	for i := 0; i < len(messages); i++ {
		creationTimestamp := messages[len(messages)-1-i].CreatedAt
		creationTimeHour := math.Floor(float64(creationTimestamp / timeGrouping))
		if timeHolding != -1 && creationTimeHour != timeHolding {
			for {
				if timeHolding+1 != creationTimeHour {
					timeCorrelationPoints = append(timeCorrelationPoints, 0)
					timeHolding++
				} else {
					break
				}
			}
			timeCorrelationPoints = append(timeCorrelationPoints, counter)
			counter = 0
		}
		counter++
		timeHolding = creationTimeHour
	}
	timeCorrelationPoints = append(timeCorrelationPoints, counter)
	return timeCorrelationPoints
}

func SeperateUsers(messages []Message) map[string][]int {
	seperatedMap := make(map[string][]int)
	seperatedMessages := make(map[string][]Message)
	for i := 0; i < len(messages); i++ {
		senderID := messages[i].Name
		seperatedMessages[senderID] = append(seperatedMessages[senderID], messages[i])
	}
	startValue := messages[len(messages)-1].CreatedAt
	for k, v := range seperatedMessages {
		output := GenerateDatapoints(v, startValue)
		seperatedMap[k] = output
	}
	return seperatedMap
}
