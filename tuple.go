package main

import "math"

type tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func clampFloat(f float64) int {
	return max(0, min(int(f), 255))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func point(x, y, z float64) tuple {
	return tuple{x, y, z, 1}
}

func vector(x, y, z float64) tuple {
	return tuple{x, y, z, 0}
}

func (t tuple) Add(x tuple) tuple {
	return tuple{t.X + x.X, t.Y + x.Y, t.Z + x.Z, t.W + x.W}
}

func (t tuple) Cross(x tuple) tuple {
	return vector(t.Y*x.Z-t.Z*x.Y, t.Z*x.X-t.X*x.Z, t.X*x.Y-t.Y*x.X)
}

func (t tuple) Div(x float64) tuple {
	return tuple{t.X / x, t.Y / x, t.Z / x, t.W}
}

func (t tuple) Dot(b tuple) float64 {
	return t.X*b.X + t.Y*b.Y + t.Z*b.Z + t.W*b.W
}

func (t tuple) Equals(x tuple) bool {
	return floatsEqual(t.X, x.X) &&
		floatsEqual(t.Y, x.Y) &&
		floatsEqual(t.Z, x.Z) &&
		floatsEqual(t.Z, x.Z)
}

func (t tuple) HadamardProduct(color tuple) tuple {
	return tuple{t.X * color.X, t.Y * color.Y, t.Z * color.Z, 0}
}

func (t tuple) Int() (int, int, int, int) {
	return clampFloat(t.X * 255), clampFloat(t.Y * 255), clampFloat(t.Z * 255), clampFloat(t.W * 255)
}

func (t tuple) IsVector() bool {
	return floatsEqual(t.W, 0)
}

func (t tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t tuple) Mul(x float64) tuple {
	return tuple{t.X * x, t.Y * x, t.Z * x, t.W}
}

func (t tuple) Negate() tuple {
	return tuple{0 - t.X, 0 - t.Y, 0 - t.Z, 0 - t.W}
}

func (t tuple) Normalize() tuple {
	m := t.Magnitude()
	return tuple{t.X / m, t.Y / m, t.Z / m, t.W / m}
}

func (t tuple) Sub(x tuple) tuple {
	return tuple{t.X - x.X, t.Y - x.Y, t.Z - x.Z, t.W - x.W}
}

func (t tuple) Translation() matrix {
	m := identity
	m[0][3] = t.X
	m[1][3] = t.Y
	m[2][3] = t.Z
	return m
}

func floatsEqual(x, y float64) bool {
	return math.Abs(x-y) < 0.001
}
