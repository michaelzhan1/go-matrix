package matrix

import "simd"

var width = simd.VectorBitSize() / 32

// this file is needed because Go will not compile if SIMD is used in a struct.
// these functions all take destination vectors to help with benchmarking.

func transposeTile(mat, dest *Matrix, rows, cols, rowTile, colTile int) {
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

func transposeRecursive(mat, dest *Matrix, rows, cols, rowTile, colTile int) {
	var helper func(i, j, rowEnd, colEnd int)
	helper = func(i, j, rowEnd, colEnd int) {
		if rowEnd-i <= rowTile && colEnd-j <= colTile {
			for k := i; k < rowEnd; k++ {
				for l := j; l < colEnd; l++ {
					dest.values[l*dest.stride+k] = mat.values[k*mat.stride+l]
				}
			}
			return
		}

		rowMid := (i + rowEnd) / 2
		colMid := (j + colEnd) / 2

		helper(i, j, rowMid, colMid)
		helper(i, colMid, rowMid, colEnd)
		helper(rowMid, j, rowEnd, colMid)
		helper(rowMid, colMid, rowEnd, colEnd)
	}

	helper(0, 0, rows, cols)
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
