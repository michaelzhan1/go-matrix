package matrix_test

import (
	"testing"

	"github.com/michaelzhan1/go-matrix/matrix/matrix"
)

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

	for i := 0; i < result.Rows(); i++ {
		for j := 0; j < result.Cols(); j++ {
			if result.Get(i, j) != expected.Get(i, j) {
				t.Errorf("Add failed at (%d,%d): got %v, want %v", i, j, result.Get(i, j), expected.Get(i, j))
			}
		}
	}
}
