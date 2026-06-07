package main

import "slices"

type ValueScore struct {
	Value int
	Score int
}

func LCV(variable string, variables map[string][]int, graphArc map[string][]string) []int {
	scores := []ValueScore{}

	for _, value := range variables[variable] {
		removed := 0

		for _, neighbor := range graphArc[variable] {
			for _, neighborValue := range variables[neighbor] {
				if !Satisfy(variable, neighbor, value, neighborValue) {
					removed++
				}

			}
		}

		scores = append(scores, ValueScore{Value: value, Score: removed})
	}

	slices.SortFunc(scores, func(i, j ValueScore) int {
		return i.Score - j.Score
	})

	result := []int{}

	for _, score := range scores {
		result = append(result, score.Value)
	}

	return result
}
