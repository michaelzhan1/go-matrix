package matrix

import "testing"

func TestMatMul(t *testing.T) {
	a := From2DSliceWithPadding([][]float32{
		{1, 2},
		{3, 4},
	})
	b := From2DSliceWithPadding([][]float32{
		{1, 2},
		{3, 4},
	})
	exp := From2DSliceWithPadding([][]float32{
		{7, 10},
		{15, 22},
	})

	out := NewMatrixWithPadding(2, 2)
	naiveMatMul(a, b, out)

	if exp.Rows() != out.Rows() || exp.Cols() != out.Cols() {
		t.Errorf("wrong output dimensions. want: %d by %d, got: %d by %d", exp.Rows(), exp.Cols(), out.Rows(), out.Cols())
	}

	for i := range out.Rows() {
		for j := range out.Cols() {
			if out.Get(i, j) != exp.Get(i, j) {
				t.Errorf("wrong value at position (%d, %d). want: %f, got: %f", i, j, exp.Get(i, j), out.Get(i, j))
			}
		}
	}
}
