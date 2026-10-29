package vector_test

import (
	"testing"

	"github.com/michaelzhan1/go-matrix/matrix/vector"
)

func TestNewRandomVector(t *testing.T) {
	tcases := []struct {
		n   int
		min float32
		max float32
	}{
		{n: 0, min: 1, max: 5},
		{n: 20, min: 12, max: 54},
		{n: 30, min: 133, max: 544},
		{n: 50, min: 22, max: 5333},
	}

	for _, tcase := range tcases {
		v := vector.NewRandomVector(tcase.n, tcase.min, tcase.max)
		if len(v) != tcase.n {
			t.Errorf("Unexpected vector length: got %v, want %v", len(v), tcase.n)
		}

		for i := range v {
			if v[i] < tcase.min || v[i] > tcase.max {
				t.Errorf("Unexpected vector value: %v is not between %v and %v", v[i], tcase.min, tcase.max)
			}
		}
	}
}
