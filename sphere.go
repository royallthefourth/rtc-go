package main

import "math"

type sphere struct {
	Origin    tuple
	Radius    float64
	Transform Matrix
}

func newSphere() sphere {
	return sphere{
		Origin:    point(0, 0, 0),
		Radius:    1,
		Transform: identity(),
	}
}

func (s sphere) Intersect(r ray) []intersection {
	// TODO apply transform
	toRay := r.Origin.Sub(s.Origin)
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(toRay)
	c := toRay.Dot(toRay) - 1
	discriminant := (b * b) - 4*a*c
	if discriminant < 0 { // that's a miss
		return []intersection{}
	}
	out := make([]intersection, 2)
	out[0].T = (-1*b - math.Sqrt(discriminant)) / (2 * a)
	out[0].Obj = s
	out[1].T = (-1*b + math.Sqrt(discriminant)) / (2 * a)
	out[1].Obj = s
	return out
}
