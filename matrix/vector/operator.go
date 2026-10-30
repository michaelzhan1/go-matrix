package vector

import "simd"

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

// scalarMul multiplies a vector by a scalar using SIMD instructions.
// It is a helper because Go fails to compile the SIMD operations in a method.
// It writes the result to the dest vector for easier benchmarking.
func scalarMul(vec Vector, scalar float32, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v := v1.Mul(simd.BroadcastFloat32s(scalar))
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v := v1.Mul(simd.BroadcastFloat32s(scalar))
		v.StorePart(dest[i:])
	}
}
