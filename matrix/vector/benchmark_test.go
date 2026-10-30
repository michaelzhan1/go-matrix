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
