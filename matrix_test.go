package main

import (
	"math"
	"testing"
)

func TestMatrix_Cofactor(t *testing.T) {
	m := matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}
	res1 := m.Cofactor(0, 0)
	res2 := m.Cofactor(1, 0)

	if !floatsEqual(res1, -12) {
		t.Errorf(`Expected %f but got %f`, float64(-12), res1)
	}

	if !floatsEqual(res2, -25) {
		t.Errorf(`Expected %f but got %f`, float64(-25), res2)
	}
}

func TestMatrix_Determinant(t *testing.T) {
	if !floatsEqual(matrix{{1, 5}, {-3, 2}}.Determinant(), 17) {
		t.Errorf(`Expected 17 but got %f`, matrix{{1, 5}, {-3, 2}}.Determinant())
	}

	tests := []struct {
		M matrix
		D float64
	}{
		{matrix{{1, 2, 6}, {-5, 8, -4}, {2, 6, 4}}, -196},
		{matrix{{-2, -8, 3, 5}, {-3, 1, 7, 3}, {1, 2, -9, 6}, {-6, 7, 7, -9}}, -4071},
	}

	for _, tst := range tests {
		if !floatsEqual(tst.M.Determinant(), tst.D) {
			t.Errorf(`Expected %f but got %f`, tst.D, tst.M.Determinant())
		}
	}
}

func TestMatrix_Equals(t *testing.T) {
	m := matrix{{1, 2}, {3, 4}}
	n := matrix{{1, 2}, {3, 4}}
	o := matrix{{1, 2}, {2, 4}}

	if !m.Equals(n) {
		t.Errorf(`%v should equal %v`, m, n)
	}

	if m.Equals(o) {
		t.Errorf(`%v should not equal %v`, m, o)
	}
}

func TestMatrix_Inverse(t *testing.T) {
	res := matrix{{-5, 2, 6, -8}, {1, -5, 1, 8}, {7, 7, -6, -7}, {1, -3, 7, 4}}.Inverse()
	ref := matrix{{.21805, .45113, .24060, -.04511},
		{-.80827, -1.45677, -.44361, .52068},
		{-.07895, -.22368, -.05263, .19737},
		{-.52256, -.81391, -.30075, .30639}}

	if !ref.Equals(res) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}
}

func TestMatrix_Minor(t *testing.T) {
	res := matrix{{3, 5, 0}, {2, -1, -7}, {6, -1, 5}}.Minor(1, 0)

	if !floatsEqual(25, res) {
		t.Errorf(`Expected %f but got %f`, float64(25), res)
	}
}

func TestMatrix_Mul(t *testing.T) {
	a := matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2}}
	b := matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8}}
	prod := matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42}}

	if !a.Mul(b).Equals(prod) {
		t.Errorf(`Expected %v but got %v`, prod, a.Mul(b))
	}
}

func TestMatrix_MulTuple(t *testing.T) {
	a := matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1}}
	b := tuple{
		X: 1,
		Y: 2,
		Z: 3,
		W: 1,
	}
	c := a.MulTuple(b)

	if !c.Equals(tuple{18, 24, 33, 1}) {
		t.Errorf(`Expected 18,24,33,1 but got %v`, c)
	}
}

func TestMatrix_Submatrix(t *testing.T) {
	res := matrix{{1, 5, 0}, {-3, 2, 7}, {0, 6, -3}}.Submatrix(0, 2)
	ref := matrix{{-3, 2}, {0, 6}}

	if !res.Equals(ref) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}

	res = matrix{{-6, 1, 1, 6}, {-8, 5, 8, 6}, {-1, 0, 8, 2}, {-7, 1, -1, 1}}.Submatrix(2, 1)
	ref = matrix{{-6, 1, 6}, {-8, 8, 6}, {-7, -1, 1}}

	if !res.Equals(ref) {
		t.Errorf(`Expected %v but got %v`, ref, res)
	}
}

func TestMatrix_Transpose(t *testing.T) {
	a := matrix{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8}}
	b := matrix{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8}}

	if !a.Transpose().Equals(b) {
		t.Errorf(`Expected %v but got %v`, b, a.Transpose())
	}
}

func TestRotX(t *testing.T) {
	halfQuarter := RotX(math.Pi / 4)
	fullQuarter := RotX(math.Pi / 2)
	p := point(0,1,0)

	if !halfQuarter.MulTuple(p).Equals(point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)){
		t.Errorf(`Expected %v but got %v`, point(0, math.Sqrt(2)/2, math.Sqrt(2)/2), halfQuarter.MulTuple(p))
	}

	if !fullQuarter.MulTuple(p).Equals(point(0, 0, 1)){
		t.Errorf(`Expected %v but got %v`, point(0, 0, 1), fullQuarter.MulTuple(p))
	}
}

func TestRotY(t *testing.T) {
	halfQuarter := RotY(math.Pi / 4)
	fullQuarter := RotY(math.Pi / 2)
	p := point(0,0,1)

	if !halfQuarter.MulTuple(p).Equals(point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)){
		t.Errorf(`Expected %v but got %v`, point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2), halfQuarter.MulTuple(p))
	}

	if !fullQuarter.MulTuple(p).Equals(point(1, 0, 0)){
		t.Errorf(`Expected %v but got %v`, point(1, 0, 0), fullQuarter.MulTuple(p))
	}
}

func TestRotZ(t *testing.T) {
	halfQuarter := RotZ(math.Pi / 4)
	fullQuarter := RotZ(math.Pi / 2)
	p := point(0,1,0)

	if !halfQuarter.MulTuple(p).Equals(point(math.Sqrt(2)/-2, math.Sqrt(2)/2, 0)){
		t.Errorf(`Expected %v but got %v`, point(math.Sqrt(2)/-2,  math.Sqrt(2)/2, 0), halfQuarter.MulTuple(p))
	}

	if !fullQuarter.MulTuple(p).Equals(point(-1, 0, 0)){
		t.Errorf(`Expected %v but got %v`, point(-1, 0, 0), fullQuarter.MulTuple(p))
	}
}
