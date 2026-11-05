package matrix

import "simd"

var width = simd.VectorBitSize() / 32

// this file is needed because Go will not compile if SIMD is used in a struct.
// these functions all take destination vectors to help with benchmarking.

func add(mat, other, dest []float32) {
	i := 0
	for ; i+width <= len(mat); i += width {
		v1 := simd.LoadFloat32s(mat[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Add(v2)
		v.Store(dest[i:])
	}

	if i < len(mat) {
		v1, _ := simd.LoadFloat32sPart(mat[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Add(v2)
		v.StorePart(dest[i:])
	}
}
