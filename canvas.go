package main

import (
	"fmt"
)

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
	pixel := make([]int, 3)
	var lineLength int

	for i := uint(0); i < c.Height; i++ {
		lineLength = 0
		for j := uint(0); j < c.Width; j++ {
			pixel[0], pixel[1], pixel[2], _ = c.Image[i][j].Int()
			for k := 0; k < 3; k++ {
				next := fmt.Sprintf(`%d`, pixel[k])
				if len(next)+lineLength+1 > 70 {
					out <- "\n"
					lineLength = len(next)
				} else if lineLength > 0 {
					out <- ` `
					lineLength++
				}
				out <- next
				lineLength += len(next)
			}
		}
		out <- "\n"
	}

	close(out)
}

func (c canvas) SetPixel(x, y int, color tuple) {
	c.Image[y][x] = color
}
