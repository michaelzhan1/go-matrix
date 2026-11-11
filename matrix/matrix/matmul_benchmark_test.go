package matrix

import (
	"testing"
)

func BenchmarkMatMul(b *testing.B) {
	funcs := []struct {
		name string
		fn   func(m1, m2, dst *Matrix)
	}{
		{"ikj", matMulIKJ},
		{"ijk", matMulIJK},
	}

	rows := 1024
	cols := 1024

	for _, f := range funcs {
		b.Run(f.name, func(b *testing.B) {
			m1 := NewRandomMatrixWithPadding(rows, cols, 0, 100)
			m2 := NewRandomMatrixWithPadding(rows, cols, 0, 100)
			dst := NewMatrixWithPadding(rows, cols)

			b.ResetTimer()
			for range b.N {
				f.fn(m1, m2, dst)
			}
		})
	}

	b.ResetTimer()
}
