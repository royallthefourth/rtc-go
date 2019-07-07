package main

import "testing"

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
