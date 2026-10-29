package vector

import "simd"

// 32-bit float vectors for now

var width = simd.VectorBitSize() / 32

type Vector []float32

// Add adds two vectors element-wise. It assumes both vectors have the same length.
func (vec Vector) Add(other Vector) Vector {
	res := make(Vector, len(vec))
	for i := range vec {
		res[i] = vec[i] + other[i]
	}
	return res
}

// AddSIMD adds two vectors element-wise using SIMD instructions.
// It assumes both vectors have the same length.
func (vec Vector) AddSIMD(other Vector) Vector {
	return addSIMD(vec, other)
}

// addSIMD adds two vectors element-wise using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
func addSIMD(vec, other Vector) Vector {
	res := make(Vector, len(vec))
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Add(v2)
		v.Store(res[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Add(v2)
		v.StorePart(res[i:])
	}

	return res
}
