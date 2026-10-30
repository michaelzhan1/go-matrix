package vector

import (
	"simd"
)

// 32-bit float vectors for now

var width = simd.VectorBitSize() / 32

type Vector []float32

// Add adds two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Add(other Vector) Vector {
	dest := make(Vector, len(vec))
	add(vec, other, dest)
	return dest
}

// add adds two vectors element-wise using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
// It writes the result to the dest vector for easier benchmarking
func add(vec, other, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Add(v2)
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Add(v2)
		v.StorePart(dest[i:])
	}
}

// Sub subtracts two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Sub(other Vector) Vector {
	dest := make(Vector, len(vec))
	sub(vec, other, dest)
	return dest
}

// sub subtracts two vectors element-wise using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
// It writes the result to the dest vector for easier benchmarking.
func sub(vec, other, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Sub(v2)
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Sub(v2)
		v.StorePart(dest[i:])
	}
}

// Mul multiplies two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Mul(other Vector) Vector {
	dest := make(Vector, len(vec))
	mul(vec, other, dest)
	return dest
}

// mul multiplies two vectors element-wise using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
// It writes the result to the dest vector for easier benchmarking.
func mul(vec, other, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Mul(v2)
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Mul(v2)
		v.StorePart(dest[i:])
	}
}

// Div divides two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Div(other Vector) Vector {
	dest := make(Vector, len(vec))
	div(vec, other, dest)
	return dest
}

// div divides two vectors element-wise using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
// It writes the result to the dest vector for easier benchmarking.
func div(vec, other, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Div(v2)
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Div(v2)
		v.StorePart(dest[i:])
	}
}
