package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestHit(t *testing.T) {
	s := newSphere()
	tests := []struct {
		Ints []intersection
		Hit  intersection
		Err  error
	}{
		{[]intersection{{1, s}, {2, s}}, intersection{1, s}, nil},
		{[]intersection{{-1, s}, {1, s}}, intersection{1, s}, nil},
		{[]intersection{{-2, s}, {-1, s}}, intersection{}, errors.New(`miss`)},
		{[]intersection{{5, s}, {7, s}, {-3, s}, {2, s}}, intersection{2, s}, nil},
	}

	for _, test := range tests {
		res, err := hit(test.Ints)
		if err != nil && err.Error() != test.Err.Error() {
			t.Errorf(`Expected error of %v but got %v`, test.Err, err)
		} else if !reflect.DeepEqual(res, test.Hit) {
			t.Errorf(`Expected hit of %v but got %v`, test.Hit, res)
		}
	}
}
