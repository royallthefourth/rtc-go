package main

import "fmt"

type canvas struct {
	Image  [][]tuple
	Height uint
	Width  uint
}

func newCanvas(width, height uint) canvas {
	c := canvas{Height: height, Width: width}
	c.Image = make([][]tuple, height)

	for i := uint(0); i < height; i++ {
		c.Image[i] = make([]tuple, width)
	}

	return c
}

func (c canvas) GetPixel(x, y int) tuple {
	return c.Image[y][x]
}

func (c canvas) Ppm(out chan string) {
	out <- fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
	close(out)
}

func (c canvas) SetPixel(x, y int, color tuple) {
	c.Image[y][x] = color
}
