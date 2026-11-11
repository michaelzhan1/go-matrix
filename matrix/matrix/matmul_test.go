package matrix

import (
	"math"
	"strconv"
	"testing"

	"simd"
)

func TestMatMulKernels(t *testing.T) {
	width := simd.VectorBitSize() / 32
	columnCounts := []int{width - 1, width, width + 1, 2 * width, 4 * width, 4*width + 1}
	kernels := []struct {
		name string
		fn   func(mat, other, dest *Matrix)
	}{
		{name: "ikj", fn: matMulIKJ},
		{name: "ijk", fn: matMulIJK},
		{name: "ijk-unrolled4", fn: matMulIJKUnrolled4},
		{name: "ijk-rows4", fn: matMulIJKRows4},
	}

	for _, cols := range columnCounts {
		t.Run(testName(cols), func(t *testing.T) {
			mat := NewMatrix(5, 5)
			other := NewMatrix(5, cols)
			for i := range mat.rows {
				for k := range mat.cols {
					mat.Set(i, k, float32((i+1)*(k+2))/3)
				}
			}
			for k := range other.rows {
				for j := range other.cols {
					other.Set(k, j, float32((k+2)*(j%7+1))/5)
				}
			}

			expected := NewMatrix(5, cols)
			naiveMatMul(mat, other, expected)

			for _, kernel := range kernels {
				t.Run(kernel.name, func(t *testing.T) {
					dest := NewMatrix(5, cols)
					kernel.fn(mat, other, dest)

					for i := range dest.rows {
						for j := range dest.cols {
							got, want := dest.Get(i, j), expected.Get(i, j)
							tolerance := 1e-5 * math.Max(1, math.Abs(float64(want)))
							if math.Abs(float64(got-want)) > tolerance {
								t.Errorf("value at (%d,%d): got %v, want %v", i, j, got, want)
							}
						}
					}
				})
			}
		})
	}
}

func testName(cols int) string {
	return "cols-" + strconv.Itoa(cols)
}
