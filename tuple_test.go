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

func TestTuple_Scaling(t *testing.T) {
	trans := point(2, 3, 4).Scaling()
	res := trans.MulTuple(point(-4, 6, 8))
	ref := point(-8, 18, 32)

	if !res.Equals(ref) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}

	res = trans.MulTuple(vector(-4, 6, 8))
	ref = vector(-8, 18, 32)
	if !res.Equals(ref) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}

	res = trans.Inverse().MulTuple(vector(-4, 6, 8))
	if !res.Equals(vector(-2, 2, 2)) {
		t.Errorf(`Expected %v but got %v`, vector(-2, 2, 2), res)
	}
}

func TestTuple_Translation(t *testing.T) {
	trans := point(5, -3, 2).Translation()
	res := trans.MulTuple(point(-3, 4, 5))
	ref := point(2, 1, 7)

	if !res.Equals(ref) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}

	res = trans.Inverse().MulTuple(point(-3, 4, 5))
	ref = point(-8, 7, 3)
	if !res.Equals(ref) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}

	res = trans.MulTuple(vector(-3, 4, 5))
	if !res.Equals(vector(-3, 4, 5)) {
		t.Errorf(`Expected %v but got %v`, vector(-3, 4, 5), res)
	}
}
