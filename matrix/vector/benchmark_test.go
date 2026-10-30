package vector

import "testing"

func BenchmarkVectorAdd(b *testing.B) {
	for _, tc := range []struct {
		name string
		fn   func(*testing.B, int)
	}{
		{"Regular", benchmarkVectorAdd},
		{"SIMD", benchmarkVectorAddSIMD},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for _, size := range []struct {
				name string
				n    int
			}{
				{"Small", 1 << 5},
				{"Medium", 1 << 10},
				{"Large", 1 << 20},
				{"NonMultiple", 999999},
			} {
				b.Run(size.name, func(b *testing.B) {
					tc.fn(b, size.n)
				})
			}
		})
	}
}

func BenchmarkVectorSub(b *testing.B) {
	for _, tc := range []struct {
		name string
		fn   func(*testing.B, int)
	}{
		{"Regular", benchmarkVectorSub},
		{"SIMD", benchmarkVectorSubSIMD},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for _, size := range []struct {
				name string
				n    int
			}{
				{"Small", 1 << 5},
				{"Medium", 1 << 10},
				{"Large", 1 << 20},
				{"NonMultiple", 999999},
			} {
				b.Run(size.name, func(b *testing.B) {
					tc.fn(b, size.n)
				})
			}
		})
	}
}

func BenchmarkVectorMul(b *testing.B) {
	for _, tc := range []struct {
		name string
		fn   func(*testing.B, int)
	}{
		{"Regular", benchmarkVectorMul},
		{"SIMD", benchmarkVectorMulSIMD},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for _, size := range []struct {
				name string
				n    int
			}{
				{"Small", 1 << 5},
				{"Medium", 1 << 10},
				{"Large", 1 << 20},
				{"NonMultiple", 999999},
			} {
				b.Run(size.name, func(b *testing.B) {
					tc.fn(b, size.n)
				})
			}
		})
	}
}

func BenchmarkVectorDiv(b *testing.B) {
	for _, tc := range []struct {
		name string
		fn   func(*testing.B, int)
	}{
		{"Regular", benchmarkVectorDiv},
		{"SIMD", benchmarkVectorDivSIMD},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for _, size := range []struct {
				name string
				n    int
			}{
				{"Small", 1 << 5},
				{"Medium", 1 << 10},
				{"Large", 1 << 20},
				{"NonMultiple", 999999},
			} {
				b.Run(size.name, func(b *testing.B) {
					tc.fn(b, size.n)
				})
			}
		})
	}
}

func BenchmarkVectorScalarMul(b *testing.B) {
	for _, tc := range []struct {
		name string
		fn   func(*testing.B, int)
	}{
		{"Regular", benchmarkVectorScalarMul},
		{"SIMD", benchmarkVectorScalarMulSIMD},
	} {
		b.Run(tc.name, func(b *testing.B) {
			for _, size := range []struct {
				name string
				n    int
			}{
				{"Small", 1 << 5},
				{"Medium", 1 << 10},
				{"Large", 1 << 20},
				{"NonMultiple", 999999},
			} {
				b.Run(size.name, func(b *testing.B) {
					tc.fn(b, size.n)
				})
			}
		})
	}
}

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
