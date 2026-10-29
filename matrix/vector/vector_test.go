package vector_test

import (
	"testing"

	"github.com/michaelzhan1/go-matrix/matrix/vector"
)

func TestAddSIMD(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	v2 := vector.Vector{4, 5, 6}
	res := v1.Add(v2)
	expected := vector.Vector{5, 7, 9}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("Add failed: got %v, want %v", res, expected)
			return
		}
	}
}
