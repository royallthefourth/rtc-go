package main

import (
	"strings"
	"testing"
)

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

func TestCanvas_Ppm(t *testing.T) {
	b := strings.Builder{}
	c := newCanvas(5, 3)
	c.SetPixel(0, 0, tuple{1.5, 0, 0, 0})
	c.SetPixel(2, 1, tuple{0, 0.5, 0, 0})
	c.SetPixel(4, 2, tuple{-0.5, 0, 1, 0})
	ch := make(chan string)
	go c.Ppm(ch)
	for line := range ch {
		_, err := b.WriteString(line)
		if err != nil {
			t.Fatal(err)
		}
	}

	rawPpm := b.String()
	lines := strings.Split(rawPpm, "\n")

	if lines[0] != `P3` || lines[2] != `255` {
		t.Errorf(`Incorrect PPM header %s`, rawPpm)
	}

	if lines[3][:7] != `255 0 0` {
		t.Errorf("PPM line 4 should start with \"255 0 0\" but got \n%s", rawPpm)
	}
}
