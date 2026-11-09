package matrix_test

import (
	"testing"

	"github.com/michaelzhan1/go-matrix/matrix/matrix"
)

func checkValues(t *testing.T, mat, exp *matrix.Matrix) {
	if mat.Rows() != exp.Rows() || mat.Cols() != exp.Cols() {
		t.Fatalf("Matrix dimensions mismatch: got %dx%d, want %dx%d", mat.Rows(), mat.Cols(), exp.Rows(), exp.Cols())
	}
	for i := range mat.Rows() {
		for j := range mat.Cols() {
			if mat.Get(i, j) != exp.Get(i, j) {
				t.Errorf("Matrix value mismatch at (%d,%d): got %v, want %v", i, j, mat.Get(i, j), exp.Get(i, j))
			}
		}
	}
}

func TestAdd(t *testing.T) {
	m1 := matrix.From2DSlice([][]float32{
		{1, 2, 3},
		{4, 5, 6},
	})
	m2 := matrix.From2DSlice([][]float32{
		{7, 8, 9},
		{10, 11, 12},
	})
	result := m1.Add(m2)
	expected := matrix.From2DSlice([][]float32{
		{8, 10, 12},
		{14, 16, 18},
	})

	checkValues(t, result, expected)
}

func TestSub(t *testing.T) {
	m1 := matrix.From2DSlice([][]float32{
		{7, 8, 9},
		{10, 11, 12},
	})
	m2 := matrix.From2DSlice([][]float32{
		{1, 2, 3},
		{4, 5, 6},
	})
	result := m1.Sub(m2)
	expected := matrix.From2DSlice([][]float32{
		{6, 6, 6},
		{6, 6, 6},
	})

	checkValues(t, result, expected)
}

func TestElemMul(t *testing.T) {
	m1 := matrix.From2DSlice([][]float32{
		{1, 2, 3},
		{4, 5, 6},
	})
	m2 := matrix.From2DSlice([][]float32{
		{7, 8, 9},
		{10, 11, 12},
	})
	result := m1.ElemMul(m2)
	expected := matrix.From2DSlice([][]float32{
		{7, 16, 27},
		{40, 55, 72},
	})

	checkValues(t, result, expected)
}

func TestElemDiv(t *testing.T) {
	m1 := matrix.From2DSlice([][]float32{
		{7, 8, 9},
		{10, 11, 12},
	})
	m2 := matrix.From2DSlice([][]float32{
		{1, 2, 3},
		{4, 5, 6},
	})
	result := m1.ElemDiv(m2)
	expected := matrix.From2DSlice([][]float32{
		{7, 4, 3},
		{2.5, 2.2, 2},
	})

	checkValues(t, result, expected)
}

func TestTranspose(t *testing.T) {
	input := matrix.From2DSlice([][]float32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	})
	expected := matrix.From2DSlice([][]float32{
		{1, 6, 11},
		{2, 7, 12},
		{3, 8, 13},
		{4, 9, 14},
		{5, 10, 15},
	})

	funcs := []struct {
		name string
		fn   func() *matrix.Matrix
	}{
		{"tile", input.TransposeTile},
		{"recursive", input.TransposeRecursive},
	}

	for _, f := range funcs {
		t.Run(f.name, func(t *testing.T) {
			result := f.fn()
			checkValues(t, result, expected)
		})
	}
}
