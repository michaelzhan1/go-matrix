package matrix

import "simd"

var width = simd.VectorBitSize() / 32

// this file is needed because Go will not compile if SIMD is used in a struct.
// these functions all take destination vectors to help with benchmarking.

func transpose(mat, dest *Matrix, rows, cols, rowTile, colTile int) {
	for i := 0; i < rows; i += rowTile {
		iEnd := min(i+rowTile, rows)

		for j := 0; j < cols; j += colTile {
			jEnd := min(j+colTile, cols)

			for k := i; k < iEnd; k++ {
				for l := j; l < jEnd; l++ {
					dest.values[l*rows+k] = mat.values[k*cols+l]
				}
			}
		}
	}
}

func transposePadded(mat, dest *Matrix, rows, cols, rowTile, colTile int) {
	for i := 0; i < rows; i += rowTile {
		iEnd := min(i+rowTile, rows)

		for j := 0; j < cols; j += colTile {
			jEnd := min(j+colTile, cols)

			for k := i; k < iEnd; k++ {
				for l := j; l < jEnd; l++ {
					dest.values[l*dest.stride+k] = mat.values[k*mat.stride+l]
				}
			}
		}
	}
}

// add adds two matrices and stores the result in the destination matrix
func add(mat, other, dest *Matrix) {
	i := 0
	for ; i+width <= len(mat.values); i += width {
		v1 := simd.LoadFloat32s(mat.values[i : i+width])
		v2 := simd.LoadFloat32s(other.values[i : i+width])
		v := v1.Add(v2)
		v.Store(dest.values[i:])
	}

	if i < len(mat.values) {
		v1, _ := simd.LoadFloat32sPart(mat.values[i:])
		v2, _ := simd.LoadFloat32sPart(other.values[i:])
		v := v1.Add(v2)
		v.StorePart(dest.values[i:])
	}
}

// sub subtracts two matrices and stores the result in the destination matrix
func sub(mat, other, dest *Matrix) {
	i := 0
	for ; i+width <= len(mat.values); i += width {
		v1 := simd.LoadFloat32s(mat.values[i : i+width])
		v2 := simd.LoadFloat32s(other.values[i : i+width])
		v := v1.Sub(v2)
		v.Store(dest.values[i:])
	}

	if i < len(mat.values) {
		v1, _ := simd.LoadFloat32sPart(mat.values[i:])
		v2, _ := simd.LoadFloat32sPart(other.values[i:])
		v := v1.Sub(v2)
		v.StorePart(dest.values[i:])
	}
}

// elemMul multiplies two matrices element-wise and stores the result in the destination matrix
func elemMul(mat, other, dest *Matrix) {
	i := 0
	for ; i+width <= len(mat.values); i += width {
		v1 := simd.LoadFloat32s(mat.values[i : i+width])
		v2 := simd.LoadFloat32s(other.values[i : i+width])
		v := v1.Mul(v2)
		v.Store(dest.values[i:])
	}

	if i < len(mat.values) {
		v1, _ := simd.LoadFloat32sPart(mat.values[i:])
		v2, _ := simd.LoadFloat32sPart(other.values[i:])
		v := v1.Mul(v2)
		v.StorePart(dest.values[i:])
	}
}

// elemDiv divides two matrices element-wise and stores the result in the destination matrix
func elemDiv(mat, other, dest *Matrix) {
	i := 0
	for ; i+width <= len(mat.values); i += width {
		v1 := simd.LoadFloat32s(mat.values[i : i+width])
		v2 := simd.LoadFloat32s(other.values[i : i+width])
		v := v1.Div(v2)
		v.Store(dest.values[i:])
	}

	if i < len(mat.values) {
		v1, _ := simd.LoadFloat32sPart(mat.values[i:])
		v2, _ := simd.LoadFloat32sPart(other.values[i:])
		v := v1.Div(v2)
		v.StorePart(dest.values[i:])
	}
}
