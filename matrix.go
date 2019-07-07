package main

type matrix [][]float64

var identity = matrix{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1}}

// newMatrix returns an initialized matrix
func newMatrix(rows int, cols int) matrix {
	m := make(matrix, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, cols)
	}
	return m
}

func (m matrix) Equals(n matrix) bool {
	if len(m) != len(n) || len(m[0]) != len(n[0]) {
		return false
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(n); j++ {
			if !floatsEqual(m[i][j], n[i][j]) {
				return false
			}
		}
	}

	return true
}

func (m matrix) Mul(n matrix) matrix {
	out := newMatrix(len(m), len(n[0]))

	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out[0]); j++ {
			out[i][j] = m[i][0]*n[0][j] +
				m[i][1]*n[1][j] +
				m[i][2]*n[2][j] +
				m[i][3]*n[3][j]
		}
	}

	return out
}

func (m matrix) MulTuple(t tuple) tuple {
	o := m.Mul(matrix{{t.X}, {t.Y}, {t.Z}, {t.W}})
	return tuple{
		X: o[0][0],
		Y: o[1][0],
		Z: o[2][0],
		W: o[3][0],
	}
}

func (m matrix) Transpose() matrix {
	out := newMatrix(len(m[0]), len(m))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			out[j][i] = m[i][j]
		}
	}

	return out
}
