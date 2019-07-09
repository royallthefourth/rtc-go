package main

import "testing"

func TestRay_Position(t *testing.T) {
	tests := []struct {
		T float64
		P tuple
	}{
		{0, point(2, 3, 4)},
		{1, point(3, 3, 4)},
		{-1, point(1, 3, 4)},
		{2.5, point(4.5, 3, 4)}}

	r := ray{point(2, 3, 4), vector(1, 0, 0)}

	for _, test := range tests {
		if !r.Position(test.T).Equals(test.P) {
			t.Errorf(`Expected %v but got %v`, test.P, r.Position(test.T))
		}
	}
}

func TestRay_Transform(t *testing.T) {
	ray1 := ray{
		Origin:    point(1, 2, 3),
		Direction: vector(0, 1, 0),
	}
	ray2 := ray1

	ray1.Transform(point(3, 4, 5).Translation())
	ray2.Transform(point(2, 3, 4).Scaling())

	if !ray1.Origin.Equals(point(4, 6, 8)) {
		t.Errorf(`Expected %v but got %v`, point(4, 6, 8), ray1.Origin)
	}
	if !ray1.Direction.Equals(point(0, 1, 0)) {
		t.Errorf(`Expected %v but got %v`, point(0, 1, 0), ray1.Direction)
	}

	if !ray2.Origin.Equals(point(2, 6, 12)) {
		t.Errorf(`Expected %v but got %v`, point(2, 6, 12), ray2.Origin)
	}
	if !ray2.Direction.Equals(point(0, 3, 0)) {
		t.Errorf(`Expected %v but got %v`, point(0, 3, 0), ray2.Direction)
	}
}
