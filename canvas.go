package main

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

func (c canvas) Ppm() string {
	out := make([]byte, 11)

}

func (c canvas) SetPixel(x, y int, color tuple) {
	c.Image[y][x] = color
}
