package main

import "testing"

func TestCanvas_GetPixel(t *testing.T) {
	c := newCanvas(10, 20)
	if !c.GetPixel(5, 5).Equals(tuple{0, 0, 0, 0}) {
		t.Errorf(`Expected black but got %+v`, c.GetPixel(5, 5))
	}

	c.SetPixel(1, 1, tuple{1, 0, 0, 0})
	if !c.GetPixel(1, 1).Equals(tuple{1, 0, 0, 0}) {
		t.Errorf(`Expected black but got %+v`, c.GetPixel(1, 1))
	}
}
