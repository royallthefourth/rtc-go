package main

type ray struct {
	Origin    tuple
	Direction tuple
}

func (r ray) Position(t float64) tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r *ray) Transform(m Matrix) {
	r.Origin = m.MulTuple(r.Origin)
	r.Direction = m.MulTuple(r.Direction)
}
