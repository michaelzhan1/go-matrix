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
