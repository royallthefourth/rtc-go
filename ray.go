package main

type ray struct {
	Origin    tuple
	Direction tuple
}

func (r ray) Position(t float64) tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}
