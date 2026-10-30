package vector

import (
	"math"
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

// ScalarAdd adds a scalar to each element of the vector using SIMD instructions.
func (vec Vector) ScalarAdd(scalar float32) Vector {
	dest := make(Vector, len(vec))
	scalarAdd(vec, scalar, dest)
	return dest
}

// ScalarSub subtracts a scalar from each element of the vector using SIMD instructions.
func (vec Vector) ScalarSub(scalar float32) Vector {
	dest := make(Vector, len(vec))
	scalarSub(vec, scalar, dest)
	return dest
}

// ScalarMul multiplies a vector by a scalar using SIMD instructions.
func (vec Vector) ScalarMul(scalar float32) Vector {
	dest := make(Vector, len(vec))
	scalarMul(vec, scalar, dest)
	return dest
}

// ScalarDiv divides a vector by a scalar using SIMD instructions.
func (vec Vector) ScalarDiv(scalar float32) Vector {
	dest := make(Vector, len(vec))
	scalarDiv(vec, scalar, dest)
	return dest
}

// Dot computes the dot product of two vectors using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) Dot(other Vector) float32 {
	return dot(vec, other)
}

// Norm computes the L2 norm of the vector.
func (vec Vector) Norm() float32 {
	return float32(math.Sqrt(float64(vec.Dot(vec))))
}

// Sum computes the sum of the elements of the vector using SIMD instructions.
func (vec Vector) Sum() float32 {
	return sum(vec)
}

// StdDev computes the population standard deviation of the elements of the vector using SIMD instructions.
// The vector cannot be empty.
func (vec Vector) StdDev() float32 {
	mean := vec.Mean()
	tmp := vec.ScalarSub(mean)
	return float32(math.Sqrt(float64(tmp.Dot(tmp)) / float64(len(vec))))
}

// Mean computes the mean of the elements of the vector using SIMD instructions.
// The vector cannot be empty.
func (vec Vector) Mean() float32 {
	return sum(vec) / float32(len(vec))
}

// Standardize will zero and scale the vector to mean 0 and standard deviation 1.
// It will fail for constant or empty vectors.
func (vec Vector) Standardize() Vector {
	mu := vec.Mean()
	sigma := vec.StdDev()
	return vec.ScalarSub(mu).ScalarDiv(sigma)
}
