package main

func InitQueue(variable string, arc []Constraint) Queue {
	queue := Queue{}

	for _, c := range arc {
		if c.Y == variable {
			queue.Enqueue(c)
		}
	}

	return queue
}

func MAC(variable string, arc []Constraint, variables map[string][]int) bool {
	queue := InitQueue(variable, arc)

	// fmt.Printf("Initial queue: %v\n", queue.Constraints)

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

		// fmt.Printf("DomainX: %v, DomainY: %v\n", domainX, domainY)
		// fmt.Printf("Dequeue constraint: %s -> %s\n", x, y)
		if Revise(&domainX, &domainY, x, y) {
			if len(domainX) == 0 {
				return false
			}

			variables[x] = domainX
			// fmt.Printf("Revise %s, new domain: %v\n", x, domainX)

			for _, c2 := range arc {
				if c2.Y == x && c2.X != y {
					queue.Enqueue(c2)
					// fmt.Printf("Enqueue constraint: %s -> %s\n", c2.X, c2.Y)
				}
			}
		}

	}

	return true
}
