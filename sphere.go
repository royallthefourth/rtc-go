package main

import "math"

type sphere struct {
	Origin tuple
	Radius float64
}

func (s sphere) Intersect(r ray) []float64 {
	toRay := r.Origin.Sub(s.Origin)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(toRay)
	c := toRay.Dot(toRay) - 1
	discriminant := (b * b) - 4*a*c
	if discriminant < 0 { // that's a miss
		return []float64{}
	}
	out := make([]float64, 2)
	out[0] = (-1*b - math.Sqrt(discriminant)) / (2 * a)
	out[1] = (-1*b + math.Sqrt(discriminant)) / (2 * a)
	return out
}
