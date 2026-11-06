package matrix

import "simd"

var width = simd.VectorBitSize() / 32

// this file is needed because Go will not compile if SIMD is used in a struct.
// these functions all take destination vectors to help with benchmarking.

func transpose(mat, dest []float32, rows, cols, rowTile, colTile int) {
	for i := 0; i < rows; i += rowTile {
		iEnd := min(i+rowTile, rows)

		for j := 0; j < cols; j += colTile {
			jEnd := min(j+colTile, cols)

			for k := i; k < iEnd; k++ {
				for l := j; l < jEnd; l++ {
					dest[l*rows+k] = mat[k*cols+l]
				}
			}
		}
	}
}

// add adds two matrices and stores the result in the destination matrix
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

// sub subtracts two matrices and stores the result in the destination matrix
func sub(mat, other, dest []float32) {
	i := 0
	for ; i+width <= len(mat); i += width {
		v1 := simd.LoadFloat32s(mat[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Sub(v2)
		v.Store(dest[i:])
	}

	if i < len(mat) {
		v1, _ := simd.LoadFloat32sPart(mat[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Sub(v2)
		v.StorePart(dest[i:])
	}
}

// elemMul multiplies two matrices element-wise and stores the result in the destination matrix
func elemMul(mat, other, dest []float32) {
	i := 0
	for ; i+width <= len(mat); i += width {
		v1 := simd.LoadFloat32s(mat[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Mul(v2)
		v.Store(dest[i:])
	}

	if i < len(mat) {
		v1, _ := simd.LoadFloat32sPart(mat[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Mul(v2)
		v.StorePart(dest[i:])
	}
}

// elemDiv divides two matrices element-wise and stores the result in the destination matrix
func elemDiv(mat, other, dest []float32) {
	i := 0
	for ; i+width <= len(mat); i += width {
		v1 := simd.LoadFloat32s(mat[i : i+width])
		v2 := simd.LoadFloat32s(other[i : i+width])
		v := v1.Div(v2)
		v.Store(dest[i:])
	}

	if i < len(mat) {
		v1, _ := simd.LoadFloat32sPart(mat[i:])
		v2, _ := simd.LoadFloat32sPart(other[i:])
		v := v1.Div(v2)
		v.StorePart(dest[i:])
	}
}
