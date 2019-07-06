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
