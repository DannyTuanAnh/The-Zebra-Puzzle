package main

func Degree(variables map[string][]int, arc map[string][]string) string {
	maxVar := ""
	maxDegree := -1

	for variable := range variables {
		degree := len(arc[variable])

		if degree > maxDegree {
			maxVar = variable
			maxDegree = degree
		}
	}

	return maxVar
}
