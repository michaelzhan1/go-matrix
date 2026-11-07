package matrix

import "math/rand/v2"

func NewRandomMatrix(rows, cols int, min, max float32) *Matrix {
	m := NewMatrix(rows, cols)
	for i := range rows {
		for j := range cols {
			m.Set(i, j, min+rand.Float32()*(max-min))
		}
	}
	return m
}

func NewRandomMatrixWithPadding(rows, cols int, min, max float32) *Matrix {
	m := NewMatrixWithPadding(rows, cols)
	for i := range rows {
		for j := range cols {
			m.Set(i, j, min+rand.Float32()*(max-min))
		}
	}
	return m
}
