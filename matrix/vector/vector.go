package vector

import (
	"simd"
)

// 32-bit float vectors for now

var width = simd.VectorBitSize() / 32

// Vector represents a 32-bit float vector.
type Vector []float32

// Add adds two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Add(other Vector) Vector {
	dest := make(Vector, len(vec))
	add(vec, other, dest)
	return dest
}

// Sub subtracts two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Sub(other Vector) Vector {
	dest := make(Vector, len(vec))
	sub(vec, other, dest)
	return dest
}

// Mul multiplies two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Mul(other Vector) Vector {
	dest := make(Vector, len(vec))
	mul(vec, other, dest)
	return dest
}

// Div divides two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Div(other Vector) Vector {
	dest := make(Vector, len(vec))
	div(vec, other, dest)
	return dest
}

// ScalarMul multiplies a vector by a scalar using SIMD instructions.
// It assumes the vector has a fixed length.
func (vec Vector) ScalarMul(scalar float32) Vector {
	dest := make(Vector, len(vec))
	scalarMul(vec, scalar, dest)
	return dest
}

// ScalarDiv divides a vector by a scalar using SIMD instructions.
// It assumes the vector has a fixed length.
func (vec Vector) ScalarDiv(scalar float32) Vector {
	dest := make(Vector, len(vec))
	scalarDiv(vec, scalar, dest)
	return dest
}
