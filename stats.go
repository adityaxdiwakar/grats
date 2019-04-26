package main

import (
	"math"
)

func GenerateDatapoints(messages []Message) []int {
	var timeCorrelationPoints []int
	timeGrouping := 3600
	counter := -1
	timeHolding := float64(-1)
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
