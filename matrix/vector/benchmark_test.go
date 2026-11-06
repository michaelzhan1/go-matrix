package vector

import "testing"

type benchCase struct {
	name string
	fn   func(*testing.B, int)
}

type benchSize struct {
	name string
	n    int
}

var benchSizes = []benchSize{
	{"Small", 1 << 6},
	{"Medium", 1 << 14},
	{"Large", 1 << 22},
	{"NonMultiple", 999999},
}

func runBenchmarks(b *testing.B, name string, fn func(*testing.B, int)) {
	b.Run(name, func(b *testing.B) {
		for _, size := range benchSizes {
			b.Run(size.name, func(b *testing.B) {
				fn(b, size.n)
			})
		}
	})
}

func BenchmarkVectorAdd(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorAdd},
		{"SIMD", benchmarkVectorAddSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorSub(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorSub},
		{"SIMD", benchmarkVectorSubSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorMul(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorMul},
		{"SIMD", benchmarkVectorMulSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorDiv(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorDiv},
		{"SIMD", benchmarkVectorDivSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorScalarAdd(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorScalarAdd},
		{"SIMD", benchmarkVectorScalarAddSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorScalarSub(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorScalarSub},
		{"SIMD", benchmarkVectorScalarSubSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorScalarMul(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorScalarMul},
		{"SIMD", benchmarkVectorScalarMulSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorScalarDiv(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorScalarDiv},
		{"SIMD", benchmarkVectorScalarDivSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorDot(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorDot},
		{"SIMD", benchmarkVectorDotSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}

func BenchmarkVectorSum(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkVectorSum},
		{"SIMD", benchmarkVectorSumSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}
