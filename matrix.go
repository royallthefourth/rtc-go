package main

import "math"

type matrix [][]float64

// newMatrix returns an initialized matrix
func newMatrix(rows, cols int) matrix {
	m := make(matrix, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]float64, cols)
	}
	return m
}

func identity() matrix {
	return matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1}}
}

func RotX(radians float64) matrix {
	m := identity()
	m[1][1] = math.Cos(radians)
	m[1][2] = -1 * math.Sin(radians)
	m[2][1] = math.Sin(radians)
	m[2][2] = math.Cos(radians)
	return m
}

func RotY(radians float64) matrix {
	m := identity()
	m[0][0] = math.Cos(radians)
	m[0][2] = math.Sin(radians)
	m[2][0] = -1 * math.Sin(radians)
	m[2][2] = math.Cos(radians)
	return m
}

func RotZ(radians float64) matrix {
	m := identity()
	m[0][0] = math.Cos(radians)
	m[0][1] = -1 * math.Sin(radians)
	m[1][0] = math.Sin(radians)
	m[1][1] = math.Cos(radians)
	return m
}

func Shear(xy, xz, yx, yz, zx, zy float64) matrix {
	m := identity()
	m[0][1] = xy
	m[0][2] = xz
	m[1][0] = yx
	m[1][2] = yz
	m[2][0] = zx
	m[2][1] = zy
	return m
}

func (m matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)

	if (row+col)%2 == 1 {
		return -1 * minor
	}

	return minor
}

func (m matrix) Determinant() float64 {
	if len(m) == 2 && len(m[0]) == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}

	det := float64(0)
	for i := 0; i < len(m[0]); i++ {
		det += m[0][i] * m.Cofactor(0, i)
	}

	return det
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

func (m matrix) Inverse() matrix {
	out := newMatrix(len(m), len(m[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			out[j][i] = m.Cofactor(i, j) / m.Determinant() // this will panic if determinant is 0, meaning matrix is not invertible
		}
	}

	return out
}

func (m matrix) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
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

func (m matrix) Submatrix(row, col int) matrix {
	out := newMatrix(len(m)-1, len(m[0])-1)

	var jOffset, iOffset int
	for i := 0; i < len(m); i++ {
		if i == row {
			continue
		}
		iOffset = i
		if i > row {
			iOffset = i - 1
		}
		for j := 0; j < len(m[0]); j++ {
			if j == col {
				continue
			}
			jOffset = j
			if j > col {
				jOffset = j - 1
			}

			out[iOffset][jOffset] = m[i][j]
		}
	}

	return out
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
