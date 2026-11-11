package matrix

import "testing"

func BenchmarkMatMul(b *testing.B) {
	funcs := []struct {
		name string
		fn   func(m1, m2, dst *Matrix)
	}{
		{name: "ikj", fn: matMulIKJ},
		{name: "ijk", fn: matMulIJK},
		{name: "ijk-unrolled4", fn: matMulIJKUnrolled4},
		{name: "ijk-rows4", fn: matMulIJKRows4},
	}

	shapes := []struct {
		name        string
		rows, inner int
		cols        int
	}{
		{name: "square-128", rows: 128, inner: 128, cols: 128},
		{name: "square-384", rows: 384, inner: 384, cols: 384},
		{name: "rectangular", rows: 256, inner: 512, cols: 768},
	}

	for _, shape := range shapes {
		b.Run(shape.name, func(b *testing.B) {
			mat := NewRandomMatrixWithPadding(shape.rows, shape.inner, 0, 1)
			other := NewRandomMatrixWithPadding(shape.inner, shape.cols, 0, 1)

			for _, f := range funcs {
				b.Run(f.name, func(b *testing.B) {
					dest := NewMatrixWithPadding(shape.rows, shape.cols)

					b.ReportAllocs()
					b.ResetTimer()
					for range b.N {
						f.fn(mat, other, dest)
						clear(dest.values)
					}
					b.ReportMetric(
						2*float64(shape.rows)*float64(shape.inner)*float64(shape.cols)*float64(b.N)/b.Elapsed().Seconds()/1e9,
						"GFLOP/s",
					)
				})
			}
		})
	}
}
