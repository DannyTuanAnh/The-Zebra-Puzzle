package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Key   string
	Value int
}

var result []map[string]int

func DeepCopy(src map[string][]int) map[string][]int {
	dst := make(map[string][]int)
	for key, value := range src {
		newValue := make([]int, len(value))
		copy(newValue, value)
		dst[key] = newValue
	}

	return dst
}

func Complete(variables map[string][]int, assignment map[string]int) bool {
	return len(variables) == len(assignment)
}

func Consistent(variable string, value int, assignment map[string]int, graphArc map[string][]string) bool {
	for _, neighbor := range graphArc[variable] {
		if neighborValue, ok := assignment[neighbor]; ok {
			if !Satisfy(variable, neighbor, value, neighborValue) {
				return false
			}
		}
	}

	return true
}

func VerifySolution(assignment map[string]int, arc []Constraint) bool {
	for _, c := range arc {
		if !Satisfy(
			c.X,
			c.Y,
			assignment[c.X],
			assignment[c.Y],
		) {
			return false
		}
	}

	return true
}

func BacktrackingSearch(variables map[string][]int, arc []Constraint, assignment map[string]int, graphArc map[string][]string) {

	if Complete(variables, assignment) {
		solution := make(map[string]int)

		for variable, value := range assignment {
			solution[variable] = value
		}

		if VerifySolution(solution, arc) {
			result = append(result, solution)
		}

		return
	}

	mrv := MRV(variables, assignment)
	variable := Degree(mrv, graphArc)

	valueList, ok := variables[variable]
	if !ok {
		return
	}

	if len(valueList) > 1 {
		valueList = LCV(variable, variables, graphArc)
	}

	for _, value := range valueList {
		if !Consistent(variable, value, assignment, graphArc) {
			continue
		}

		assignment[variable] = value

		variablesCopy := DeepCopy(variables)

		variablesCopy[variable] = []int{value}

		if MAC(variable, arc, variablesCopy) {
			BacktrackingSearch(variablesCopy, arc, assignment, graphArc)
		}

		delete(assignment, variable)
	}
}

func genArc(arc *[]Constraint, attribute []string) {
	for i := 0; i < len(attribute); i++ {
		for j := 0; j < len(attribute); j++ {
			if i != j {
				*arc = append(*arc, Constraint{X: attribute[i], Y: attribute[j]})
			}
		}
	}
}

func main() {
	variables := map[string][]int{
		"Yellow": {1, 2, 3, 4, 5},
		"Blue":   {1, 2, 3, 4, 5},
		"Red":    {1, 2, 3, 4, 5},
		"Green":  {1, 2, 3, 4, 5},
		"Ivory":  {1, 2, 3, 4, 5},

		"Zebra":  {1, 2, 3, 4, 5},
		"Dog":    {1, 2, 3, 4, 5},
		"Snails": {1, 2, 3, 4, 5},
		"Fox":    {1, 2, 3, 4, 5},
		"Horse":  {1, 2, 3, 4, 5},

		"Coffee":      {1, 2, 3, 4, 5},
		"Tea":         {1, 2, 3, 4, 5},
		"Milk":        {1, 2, 3, 4, 5},
		"OrangeJuice": {1, 2, 3, 4, 5},
		"Water":       {1, 2, 3, 4, 5},

		"OldGold":       {1, 2, 3, 4, 5},
		"Kools":         {1, 2, 3, 4, 5},
		"Chesterfields": {1, 2, 3, 4, 5},
		"Parliaments":   {1, 2, 3, 4, 5},
		"LuckyStrike":   {1, 2, 3, 4, 5},

		"Britt":     {1, 2, 3, 4, 5},
		"Spaniard":  {1, 2, 3, 4, 5},
		"Ukrainian": {1, 2, 3, 4, 5},
		"Norwegian": {1, 2, 3, 4, 5},
		"Japanese":  {1, 2, 3, 4, 5},
	}

	arc := []Constraint{
		{"Britt", "Red"},
		{"Red", "Britt"},

		{"Spaniard", "Dog"},
		{"Dog", "Spaniard"},

		{"Coffee", "Green"},
		{"Green", "Coffee"},

		{"Ukrainian", "Tea"},
		{"Tea", "Ukrainian"},

		{"Green", "Ivory"},
		{"Ivory", "Green"},

		{"OldGold", "Snails"},
		{"Snails", "OldGold"},

		{"Chesterfields", "Fox"},
		{"Fox", "Chesterfields"},

		{"Kools", "Horse"},
		{"Horse", "Kools"},

		{"Kools", "Yellow"},
		{"Yellow", "Kools"},

		{"LuckyStrike", "OrangeJuice"},
		{"OrangeJuice", "LuckyStrike"},

		{"Parliaments", "Japanese"},
		{"Japanese", "Parliaments"},

		{"Water", "Chesterfields"},
		{"Chesterfields", "Water"},

		{"Norwegian", "Blue"},
		{"Blue", "Norwegian"},
	}

	colors := []string{"Yellow", "Blue", "Red", "Green", "Ivory"}
	nationalities := []string{"Britt", "Spaniard", "Ukrainian", "Norwegian", "Japanese"}
	pets := []string{"Zebra", "Dog", "Snails", "Fox", "Horse"}
	drinks := []string{"Coffee", "Tea", "Milk", "OrangeJuice", "Water"}
	tobaccos := []string{"OldGold", "Kools", "Chesterfields", "Parliaments", "LuckyStrike"}

	genArc(&arc, colors)
	genArc(&arc, nationalities)
	genArc(&arc, pets)
	genArc(&arc, drinks)
	genArc(&arc, tobaccos)

	graphArc := map[string][]string{
		"Britt": {"Red", "Norwegian", "Spaniard", "Ukrainian", "Japanese"},
		"Red":   {"Britt", "Yellow", "Blue", "Green", "Ivory"},

		"Spaniard": {"Dog", "Norwegian", "Britt", "Ukrainian", "Japanese"},
		"Dog":      {"Spaniard", "Zebra", "Snails", "Fox", "Horse"},

		"Coffee": {"Green", "Tea", "Milk", "OrangeJuice", "Water"},
		"Green":  {"Coffee", "Ivory", "Yellow", "Blue", "Red"},

		"Ukrainian": {"Tea", "Norwegian", "Britt", "Spaniard", "Japanese"},
		"Tea":       {"Ukrainian", "Coffee", "Milk", "OrangeJuice", "Water"},

		"Ivory": {"Green", "Yellow", "Blue", "Red"},

		"OldGold": {"Snails", "Kools", "Chesterfields", "Parliaments", "LuckyStrike"},
		"Snails":  {"OldGold", "Zebra", "Dog", "Fox", "Horse"},

		"Chesterfields": {"Fox", "Water", "OldGold", "Kools", "Parliaments", "LuckyStrike"},
		"Fox":           {"Chesterfields", "Zebra", "Dog", "Snails", "Horse"},

		"Kools": {"Horse", "OldGold", "Chesterfields", "Parliaments", "LuckyStrike", "Yellow"},
		"Horse": {"Kools", "Zebra", "Dog", "Snails", "Fox"},

		"LuckyStrike": {"OrangeJuice", "OldGold", "Kools", "Chesterfields", "Parliaments"},
		"OrangeJuice": {"LuckyStrike", "Coffee", "Tea", "Milk", "Water"},

		"Parliaments": {"Japanese", "OldGold", "Kools", "Chesterfields", "LuckyStrike"},
		"Japanese":    {"Parliaments", "Britt", "Spaniard", "Ukrainian", "Norwegian"},

		"Norwegian": {"Blue", "Britt", "Spaniard", "Ukrainian", "Japanese"},
		"Blue":      {"Norwegian", "Yellow", "Red", "Green", "Ivory"},

		"Yellow": {"Blue", "Red", "Green", "Ivory", "Kools"},
	}

	assignment := map[string]int{}

	if AC3(arc, variables) {
		fmt.Println("Domain of each variable after using AC3:")
		fmt.Println(variables)
		fmt.Println("\nUsing MAC in backtrack")
		BacktrackingSearch(variables, arc, assignment, graphArc)
	}

	fmt.Println("Result:")
	answer := make(map[string]int)
	items := make([]Item, 0)
	for _, solution := range result {
		for key, value := range solution {
			if key == "Zebra" || key == "Water" {
				answer[key] = value
			}
			items = append(items, Item{Key: key, Value: value})
		}

		sort.Slice(items, func(i, j int) bool {
			return items[i].Value < items[j].Value
		})
	}

	count := 0
	for _, item := range items {
		count++
		fmt.Print(item.Key, ":", item.Value)
		if count%5 == 0 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Printf("\nThe owner of the zebra lives in house %d.\n", answer["Zebra"])
	fmt.Printf("The owner who drinks water lives in house %d.\n", answer["Water"])

	fmt.Println("\n\nPress Enter to exit...")
	fmt.Scanln()
}
