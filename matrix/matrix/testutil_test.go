package matrix

import "testing"

func TestNewRandomMatrix(t *testing.T) {
	tcases := []struct {
		rows int
		cols int
		min  float32
		max  float32
	}{
		{rows: 0, cols: 0, min: 1, max: 5},
		{rows: 20, cols: 10, min: 12, max: 54},
		{rows: 30, cols: 15, min: 133, max: 544},
		{rows: 50, cols: 25, min: 22, max: 5333},
	}

	for _, tcase := range tcases {
		m := NewRandomMatrix(tcase.rows, tcase.cols, tcase.min, tcase.max)
		if m.Rows() != tcase.rows || m.Cols() != tcase.cols {
			t.Errorf("Unexpected matrix dimensions: got %dx%d, want %dx%d", m.Rows(), m.Cols(), tcase.rows, tcase.cols)
		}

		for i := range m.Rows() {
			for j := range m.Cols() {
				val := m.Get(i, j)
				if val < tcase.min || val > tcase.max {
					t.Errorf("Unexpected matrix value at (%d,%d): %v is not between %v and %v", i, j, val, tcase.min, tcase.max)
				}
			}
		}
	}
}
