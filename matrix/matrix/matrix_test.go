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

	for i := range result.Rows() {
		for j := range result.Cols() {
			if result.Get(i, j) != expected.Get(i, j) {
				t.Errorf("Add failed at (%d,%d): got %v, want %v", i, j, result.Get(i, j), expected.Get(i, j))
			}
		}
	}
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

	for i := range result.Rows() {
		for j := range result.Cols() {
			if result.Get(i, j) != expected.Get(i, j) {
				t.Errorf("Sub failed at (%d,%d): got %v, want %v", i, j, result.Get(i, j), expected.Get(i, j))
			}
		}
	}
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

	for i := range result.Rows() {
		for j := range result.Cols() {
			if result.Get(i, j) != expected.Get(i, j) {
				t.Errorf("ElemMul failed at (%d,%d): got %v, want %v", i, j, result.Get(i, j), expected.Get(i, j))
			}
		}
	}
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

	for i := range result.Rows() {
		for j := range result.Cols() {
			if result.Get(i, j) != expected.Get(i, j) {
				t.Errorf("ElemDiv failed at (%d,%d): got %v, want %v", i, j, result.Get(i, j), expected.Get(i, j))
			}
		}
	}
}

func TestTranspose(t *testing.T) {
	input := matrix.From2DSlice([][]float32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	})
	result := input.Transpose()
	expected := [][]float32{
		{1, 6, 11},
		{2, 7, 12},
		{3, 8, 13},
		{4, 9, 14},
		{5, 10, 15},
	}

	if result.Rows() != len(expected) || result.Cols() != len(expected[0]) {
		t.Fatalf("Transpose dimensions = %dx%d, want %dx%d", result.Rows(), result.Cols(), len(expected), len(expected[0]))
	}
	for i, row := range expected {
		for j, want := range row {
			if got := result.Get(i, j); got != want {
				t.Errorf("Transpose value at (%d,%d) = %v, want %v", i, j, got, want)
			}
		}
	}
}
