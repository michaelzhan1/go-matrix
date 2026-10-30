package vector

import "testing"

func benchmarkVectorAdd(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] + v2[i]
		}
	}
}

func benchmarkVectorAddSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		add(v1, v2, res)
	}
}

func benchmarkVectorSub(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] - v2[i]
		}
	}
}

func benchmarkVectorSubSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		sub(v1, v2, res)
	}
}

func benchmarkVectorMul(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] * v2[i]
		}
	}
}

func benchmarkVectorMulSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		mul(v1, v2, res)
	}
}

func benchmarkVectorDiv(b *testing.B, size int) {
	v1 := NewRandomVector(size, 1, 100)
	v2 := NewRandomVector(size, 1, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] / v2[i]
		}
	}
}

func benchmarkVectorDivSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 1, 100)
	v2 := NewRandomVector(size, 1, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		div(v1, v2, res)
	}
}

func benchmarkVectorScalarMul(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] * scalar
		}
	}
}

func benchmarkVectorScalarMulSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		scalarMul(v1, scalar, res)
	}
}

func benchmarkVectorScalarDiv(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] / scalar
		}
	}
}

func benchmarkVectorScalarDivSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		scalarDiv(v1, scalar, res)
	}
}
