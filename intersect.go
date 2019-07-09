package main

import "errors"

type intersectable interface {
	Intersect(r ray) []intersection
}

type intersection struct {
	T   float64
	Obj intersectable
}

// hit returns the lowest nonnegative intersection
func hit(is []intersection) (intersection, error) {
	min := is[0]

	for _, i := range is {
		if i.T >= 0 && (i.T < min.T || min.T < 0) {
			min = i
		}
	}

	if min.T >= 0 {
		return min, nil
	}

	return intersection{}, errors.New(`miss`)
}
