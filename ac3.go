package main

import "math"

func sameGroup(x, y string) bool {
	groups := [][]string{
		{"Red", "Green", "Ivory", "Yellow", "Blue"},
		{"Britt", "Spaniard", "Ukrainian", "Norwegian", "Japanese"},
		{"Coffee", "Tea", "Milk", "OrangeJuice", "Water"},
		{"OldGold", "Kools", "Chesterfields", "Parliaments", "LuckyStrike"},
		{"Dog", "Snails", "Fox", "Horse", "Zebra"},
	}

	for _, group := range groups {
		foundX := false
		foundY := false

		for _, v := range group {
			if v == x {
				foundX = true
			}

			if v == y {
				foundY = true
			}
		}

		if foundX && foundY {
			return true
		}
	}

	return false
}

func Satisfy(x, y string, vx, vy int) bool {
	if (x == "Britt" && y == "Red") || (x == "Red" && y == "Britt") {
		return vx == vy
	}

	if (x == "Spaniard" && y == "Dog") || (x == "Dog" && y == "Spaniard") {
		return vx == vy
	}

	if (x == "Coffee" && y == "Green") || (x == "Green" && y == "Coffee") {
		return vx == vy
	}

	if (x == "Ukrainian" && y == "Tea") || (x == "Tea" && y == "Ukrainian") {
		return vx == vy
	}

	if x == "Green" && y == "Ivory" {
		return vx == vy+1
	}

	if x == "Ivory" && y == "Green" {
		return vx+1 == vy
	}

	if (x == "OldGold" && y == "Snails") || (x == "Snails" && y == "OldGold") {
		return vx == vy
	}

	if (x == "Kools" && y == "Yellow") || (x == "Yellow" && y == "Kools") {
		return vx == vy
	}

	if (x == "Chesterfields" && y == "Fox") || (x == "Fox" && y == "Chesterfields") {
		return vx == vy+1 || vx == vy-1
	}

	if (x == "Kools" && y == "Horse") || (x == "Horse" && y == "Kools") {
		return vx == vy+1 || vx == vy-1
	}

	if (x == "LuckyStrike" && y == "OrangeJuice") || (x == "OrangeJuice" && y == "LuckyStrike") {
		return vx == vy
	}

	if (x == "Parliaments" && y == "Japanese") || (x == "Japanese" && y == "Parliaments") {
		return vx == vy
	}

	if (x == "Water" && y == "Chesterfields") || (x == "Chesterfields" && y == "Water") {
		return vx == vy+1 || vx == vy-1
	}

	if (x == "Norwegian" && y == "Blue") || (x == "Blue" && y == "Norwegian") {
		return math.Abs(float64(vx-vy)) == 1
	}

	if sameGroup(x, y) {
		return vx != vy
	}

	return true
}

func Revise(x, y *[]int, xName, yName string) bool {
	revised := false

	newDomain := []int{}

	for _, vx := range *x {
		satisfied := false
		for _, vy := range *y {
			if Satisfy(xName, yName, vx, vy) {
				satisfied = true
				break
			}

		}

		if satisfied {
			newDomain = append(newDomain, vx)
		} else {
			revised = true
		}
	}

	*x = newDomain

	return revised
}

func AC3(arc []Constraint, variables map[string][]int) bool {

	queue := Queue{Constraints: arc}

	variables["Norwegian"] = []int{1}
	variables["Milk"] = []int{3}

	for len(queue.Constraints) > 0 {
		q, ok := queue.Dequeue()
		if !ok {
			return false
		}

		x := q.X
		y := q.Y

		domainX, ok := variables[x]
		if !ok {
			return false
		}

		domainY, ok := variables[y]
		if !ok {
			return false
		}

		if Revise(&domainX, &domainY, x, y) {
			if len(domainX) == 0 {
				return false
			}

			variables[x] = domainX

			for _, c2 := range arc {
				if c2.Y == x && c2.X != y {
					queue.Enqueue(c2)
				}
			}
		}

	}

	return true
}
