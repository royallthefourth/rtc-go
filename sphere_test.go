package main

import "testing"

func TestSphere_Intersect(t *testing.T) {
	tests := []struct {
		R  ray
		S  sphere
		Ts []float64
	}{
		{
			R:  ray{point(0, 0, -5), vector(0, 0, 1)},
			S:  newSphere(),
			Ts: []float64{4, 6},
		},
		{
			R:  ray{point(0, 1, -5), vector(0, 0, 1)},
			S:  newSphere(),
			Ts: []float64{5, 5},
		},
		{
			R:  ray{point(0, 2, -5), vector(0, 0, 1)},
			S:  newSphere(),
			Ts: []float64{},
		},
		{
			R:  ray{point(0, 0, 0), vector(0, 0, 1)},
			S:  newSphere(),
			Ts: []float64{-1, 1},
		},
	}

	for _, test := range tests {
		res := test.S.Intersect(test.R)
		if len(res) != len(test.Ts) {
			t.Errorf(`Expected %v but got %v`, test.Ts, res)
		}
		for ix, tVal := range test.Ts {
			if res[ix].T != tVal {
				t.Errorf(`Expected %f but got %f`, tVal, res[ix])
			}
		}
	}
}

func TestSphere_IntersectScale(t *testing.T) {
	r := ray{
		Origin:    point(0,0,-5),
		Direction: vector(0,0,1),
	}
}
