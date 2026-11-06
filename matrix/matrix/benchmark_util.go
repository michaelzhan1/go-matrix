package matrix

import "testing"

func benchmarkMatrixAdd(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 0, 100)
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		naiveAdd(m1.values, m2.values, res.values)
	}
}

func benchmarkMatrixAddSIMD(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 0, 100)
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		add(m1.values, m2.values, res.values)
	}
}
