package main

func MRVOriginal(variables map[string][]int) string {
	minVar := ""

	for variable, domain := range variables {
		if minVar == "" || len(domain) < len(variables[minVar]) {
			minVar = variable
		}
	}

	return minVar
}

func MRV(variables map[string][]int, assignment map[string]int) map[string][]int {
	minList := map[string][]int{}
	currentLen := 0

	for variable, domain := range variables {
		if _, assigned := assignment[variable]; assigned {
			continue
		}

		if currentLen == 0 {
			currentLen = len(domain)
			minList[variable] = domain
			continue
		}
		if len(domain) == currentLen {
			minList[variable] = domain
		}

		if len(domain) < currentLen {
			currentLen = len(domain)
			minList = map[string][]int{variable: domain}
		}
	}

	return minList
}
