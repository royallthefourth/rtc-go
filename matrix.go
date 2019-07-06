package main

type matrix [][]float64

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
