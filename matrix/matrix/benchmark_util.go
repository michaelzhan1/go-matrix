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
		add(m1, m2, res)
	}
}

func benchmarkMatrixSub(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 0, 100)
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		naiveSub(m1.values, m2.values, res.values)
	}
}

func benchmarkMatrixSubSIMD(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 0, 100)
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		sub(m1, m2, res)
	}
}

func benchmarkMatrixElemMul(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 0, 100)
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		naiveElemMul(m1.values, m2.values, res.values)
	}
}

func benchmarkMatrixElemMulSIMD(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 0, 100)
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		elemMul(m1, m2, res)
	}
}

func benchmarkMatrixElemDiv(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 1, 100) // Avoid division by zero
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		naiveElemDiv(m1.values, m2.values, res.values)
	}
}

func benchmarkMatrixElemDivSIMD(b *testing.B, rows, cols int) {
	m1 := NewRandomMatrix(rows, cols, 0, 100)
	m2 := NewRandomMatrix(rows, cols, 1, 100) // Avoid division by zero
	res := NewMatrix(rows, cols)

	b.ResetTimer()
	for b.Loop() {
		elemDiv(m1, m2, res)
	}
}
