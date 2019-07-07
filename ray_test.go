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
