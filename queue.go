package main

type Constraint struct {
	X string
	Y string
}

type Queue struct {
	Constraints []Constraint
}

func (q *Queue) Enqueue(c Constraint) {
	q.Constraints = append(q.Constraints, c)
}

func (q *Queue) Dequeue() (Constraint, bool) {
	if len(q.Constraints) == 0 {
		return Constraint{}, false
	}

	c := q.Constraints[0]
	q.Constraints = q.Constraints[1:]
	return c, true
}
