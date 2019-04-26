package main

import (
	"math"
)

func GenerateDatapoints(messages []Message) []int {
	var TimeCorrelationPoints []int
	TimeGrouping := 3600 * 24
	Counter := -1
	TimeHolding := float64(-1)
	for i := 0; i < len(messages); i++ {
		CreationTimestamp := messages[len(messages)-1-i].CreatedAt
		CreationTimeHour := math.Floor(float64(CreationTimestamp / TimeGrouping))
		if TimeHolding != -1 && CreationTimeHour != TimeHolding {
			for {
				if TimeHolding+1 != CreationTimeHour {
					TimeCorrelationPoints = append(TimeCorrelationPoints, 0)
					TimeHolding++
				} else {
					break
				}
			}
			TimeCorrelationPoints = append(TimeCorrelationPoints, Counter)
			Counter = 0
		}
		Counter++
		TimeHolding = CreationTimeHour
	}
	TimeCorrelationPoints = append(TimeCorrelationPoints, Counter)
	return TimeCorrelationPoints
}
