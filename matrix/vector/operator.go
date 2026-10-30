package vector

import "simd"

// this file is needed because Go will not compile if SIMD is used in a struct.
// these functions all take destination vectors to help with benchmarking.

// add adds two vectors element-wise using SIMD instructions.
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

// scalarAdd adds a scalar to each element of the vector using SIMD instructions.
func scalarAdd(vec Vector, scalar float32, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v := v1.Add(simd.BroadcastFloat32s(scalar))
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v := v1.Add(simd.BroadcastFloat32s(scalar))
		v.StorePart(dest[i:])
	}
}

// scalarSub subtracts a scalar from each element of the vector using SIMD instructions.
func scalarSub(vec Vector, scalar float32, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v := v1.Sub(simd.BroadcastFloat32s(scalar))
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v := v1.Sub(simd.BroadcastFloat32s(scalar))
		v.StorePart(dest[i:])
	}
}

// scalarMul multiplies a vector by a scalar using SIMD instructions.
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

// scalarDiv divides a vector by a scalar using SIMD instructions.
func scalarDiv(vec Vector, scalar float32, dest Vector) {
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v := v1.Div(simd.BroadcastFloat32s(scalar))
		v.Store(dest[i:])
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v := v1.Div(simd.BroadcastFloat32s(scalar))
		v.StorePart(dest[i:])
	}
}

// dot computes the dot product of two vectors using SIMD instructions.
func dot(vec, other Vector) float32 {
	v := simd.BroadcastFloat32s(0)
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v = v.Add(v1.Mul(v2))
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v = v.Add(v1.Mul(v2))
	}

	vSlice := make([]float32, width)
	v.Store(vSlice)
	var sum float32
	for _, x := range vSlice {
		sum += x
	}
	return sum
}

// sum computes the sum of the elements of a vector using SIMD instructions.
func sum(vec Vector) float32 {
	v := simd.BroadcastFloat32s(0)
	i := 0
	for ; i+width <= len(vec); i += width {
		v1 := simd.LoadFloat32s(vec[i : i+width])
		v = v.Add(v1)
	}

	if i < len(vec) {
		v1, _ := simd.LoadFloat32sPart(vec[i:])
		v = v.Add(v1)
	}

	vSlice := make([]float32, width)
	v.Store(vSlice)
	var sum float32
	for _, x := range vSlice {
		sum += x
	}
	return sum
}
