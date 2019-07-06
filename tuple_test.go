package main

import (
	"math"
	"testing"
)

func TestClampFloat(t *testing.T) {
	if clampFloat(-1) != 0 {
		t.Errorf(`Expected -1 to be 0 but got %d`, clampFloat(-1))
	}
}

func TestTuple_Magnitude(t *testing.T) {
	m := vector(1, 0, 0).Magnitude()
	if !floatsEqual(1, m) {
		t.Errorf(`Expected 1 but got %f`, m)
	}

	m = vector(1, 2, 3).Magnitude()
	sqrt14 := math.Sqrt(14)
	if !floatsEqual(sqrt14, m) {
		t.Errorf(`Expected %f but got %f`, sqrt14, m)
	}
}

func TestTuple_Normalize(t *testing.T) {
	if !vector(4, 0, 0).Normalize().Equals(vector(1, 0, 0)) {
		t.Fail()
	}
}

func TestTuple_Dot(t *testing.T) {
	if !floatsEqual(vector(1, 2, 3).Dot(vector(2, 3, 4)), 20) {
		t.Fail()
	}
}

func TestTuple_Cross(t *testing.T) {
	if !vector(1, 2, 3).Cross(vector(2, 3, 4)).Equals(vector(-1, 2, -1)) {
		t.Errorf(`Expected 1,2,3 X 2,3,4 to be -1,2,-1 but got %v`, vector(1, 2, 3).Cross(vector(2, 3, 4)))
	}

	if !vector(2, 3, 4).Cross(vector(1, 2, 3)).Equals(vector(1, -2, 1)) {
		t.Errorf(`Expected 2,3,4 X 1,2,3 to be 1,-2,1 but got %v`, vector(2, 3, 4).Cross(vector(1, 2, 3)))
	}
}
