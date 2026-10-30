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

func TestSubSIMD(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	v2 := vector.Vector{4, 5, 6}
	res := v1.Sub(v2)
	expected := vector.Vector{-3, -3, -3}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("Sub failed: got %v, want %v", res, expected)
			return
		}
	}
}

func TestMulSIMD(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	v2 := vector.Vector{4, 5, 6}
	res := v1.Mul(v2)
	expected := vector.Vector{4, 10, 18}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("Mul failed: got %v, want %v", res, expected)
			return
		}
	}
}

func TestDivSIMD(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	v2 := vector.Vector{4, 5, 6}
	res := v1.Div(v2)
	expected := vector.Vector{0.25, 0.4, 0.5}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("Div failed: got %v, want %v", res, expected)
			return
		}
	}
}
