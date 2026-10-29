package vector

import (
	"simd"
)

// 32-bit float vectors for now

var width = simd.VectorBitSize() / 32

type Vector []float32

// Add adds two vectors element-wise. It assumes both vectors have the same length.
func (vec Vector) Add(other Vector) Vector {
	res := make(Vector, len(vec))
	add(vec, other, res)
	return res
}

// AddSIMD adds two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) AddSIMD(other Vector) Vector {
	dest := make(Vector, len(vec))
	addSIMD(vec, other, dest)
	return dest
}

// add adds two vectors element-wise. It writes the result to the dest vector.
func add(vec, other, dest Vector) {
	for i := range vec {
		dest[i] = vec[i] + other[i]
	}
}

// addSIMD adds two vectors element-wise using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
// It writes the result to the dest vector for easier benchmarking
func addSIMD(vec, other, dest Vector) {
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
