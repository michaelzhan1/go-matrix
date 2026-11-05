package vector

import "testing"

func benchmarkVectorAdd(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		naiveAdd(v1, v2, res)
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
		naiveSub(v1, v2, res)
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
		naiveMul(v1, v2, res)
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
		naiveDiv(v1, v2, res)
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

func benchmarkVectorScalarAdd(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		naiveScalarAdd(v1, scalar, res)
	}
}

func benchmarkVectorScalarAddSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		scalarAdd(v1, scalar, res)
	}
}

func benchmarkVectorScalarSub(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		naiveScalarSub(v1, scalar, res)
	}
}

func benchmarkVectorScalarSubSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		scalarSub(v1, scalar, res)
	}
}

func benchmarkVectorScalarMul(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	scalar := float32(2)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		naiveScalarMul(v1, scalar, res)
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
		naiveScalarDiv(v1, scalar, res)
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

func benchmarkVectorDot(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)

	b.ResetTimer()
	for b.Loop() {
		_ = naiveDot(v1, v2)
	}
}

func benchmarkVectorDotSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)

	b.ResetTimer()
	for b.Loop() {
		_ = dot(v1, v2)
	}
}

func benchmarkVectorSum(b *testing.B, size int) {
	v := NewRandomVector(size, 0, 100)

	b.ResetTimer()
	for b.Loop() {
		_ = naiveSum(v)
	}
}

func benchmarkVectorSumSIMD(b *testing.B, size int) {
	v := NewRandomVector(size, 0, 100)

	b.ResetTimer()
	for b.Loop() {
		_ = sum(v)
	}
}
