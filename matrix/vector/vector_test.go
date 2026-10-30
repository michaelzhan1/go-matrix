package vector_test

import (
	"testing"

	"github.com/michaelzhan1/go-matrix/matrix/vector"
)

func TestAdd(t *testing.T) {
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

func TestSub(t *testing.T) {
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

func TestMul(t *testing.T) {
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

func TestDiv(t *testing.T) {
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

func TestScalarAdd(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	scalar := float32(2)
	res := v1.ScalarAdd(scalar)
	expected := vector.Vector{3, 4, 5}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("ScalarAdd failed: got %v, want %v", res, expected)
			return
		}
	}
}

func TestScalarSub(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	scalar := float32(2)
	res := v1.ScalarSub(scalar)
	expected := vector.Vector{-1, 0, 1}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("ScalarSub failed: got %v, want %v", res, expected)
			return
		}
	}
}

func TestScalarMul(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	scalar := float32(2)
	res := v1.ScalarMul(scalar)
	expected := vector.Vector{2, 4, 6}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("ScalarMul failed: got %v, want %v", res, expected)
			return
		}
	}
}

func TestScalarDiv(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	scalar := float32(2)
	res := v1.ScalarDiv(scalar)
	expected := vector.Vector{0.5, 1, 1.5}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("ScalarDiv failed: got %v, want %v", res, expected)
			return
		}
	}
}

func TestDot(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	v2 := vector.Vector{4, 5, 6}
	res := v1.Dot(v2)
	expected := float32(32)
	if res != expected {
		t.Errorf("Dot failed: got %v, want %v", res, expected)
	}
}

func TestNorm(t *testing.T) {
	v1 := vector.Vector{2, 4, 4}
	res := v1.Norm()
	expected := float32(6)
	if res != expected {
		t.Errorf("Norm failed: got %v, want %v", res, expected)
	}
}

func TestSum(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	res := v1.Sum()
	expected := float32(6)
	if res != expected {
		t.Errorf("Sum failed: got %v, want %v", res, expected)
	}
}

func TestStdDev(t *testing.T) {
	v1 := vector.Vector{1, 1, 5, 5}
	res := v1.StdDev()
	expected := float32(2)
	if res != expected {
		t.Errorf("StdDev failed: got %v, want %v", res, expected)
	}
}

func TestMean(t *testing.T) {
	v1 := vector.Vector{1, 2, 3}
	res := v1.Mean()
	expected := float32(2)
	if res != expected {
		t.Errorf("Mean failed: got %v, want %v", res, expected)
	}
}

func TestStandardize(t *testing.T) {
	v1 := vector.Vector{1, 1, 5, 5}
	res := v1.Standardize()
	expected := vector.Vector{-1, -1, 1, 1}
	for i := range res {
		if res[i] != expected[i] {
			t.Errorf("Standardize failed: got %v, want %v", res, expected)
			return
		}
	}
}
