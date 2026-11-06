package matrix

import "testing"

type benchCase struct {
	name string
	fn   func(*testing.B, int, int)
}

type benchSize struct {
	name string
	rows int
	cols int
}

var benchSizes = []benchSize{
	{"Small", 1 << 3, 1 << 3},
	{"Medium", 1 << 7, 1 << 7},
	{"Large", 1 << 11, 1 << 11},
	{"NonMultiple", 999, 999},
}

func runBenchmarks(b *testing.B, name string, fn func(*testing.B, int, int)) {
	b.Run(name, func(b *testing.B) {
		for _, size := range benchSizes {
			b.Run(size.name, func(b *testing.B) {
				fn(b, size.rows, size.cols)
			})
		}
	})
}

func BenchmarkMatrixAdd(b *testing.B) {
	for _, tc := range []benchCase{
		{"Regular", benchmarkMatrixAdd},
		{"SIMD", benchmarkMatrixAddSIMD},
	} {
		runBenchmarks(b, tc.name, tc.fn)
	}
}
