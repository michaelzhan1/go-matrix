package matrix

import (
	"simd"
	"simd/archsimd"
)

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

// transposeSIMD recursively loads tiles to transpose using SIMD instructions.
func transposeSIMD(mat, dest *Matrix) {
	var helper func(i, j, rowEnd, colEnd int)
	helper = func(i, j, rowEnd, colEnd int) {
		rowLen := rowEnd - i
		colLen := colEnd - j

		// Base case threshold: when sub-matrices are <= 8x8
		if rowLen <= 8 && colLen <= 8 {
			if rowLen == 8 && colLen == 8 {
				transposeSIMD8x8(mat, dest, i, j)
			} else {
				// Scalar fallback for odd boundary remainder tails (< 8 rows or cols)
				for k := i; k < rowEnd; k++ {
					srcRowOffset := k * mat.stride
					for l := j; l < colEnd; l++ {
						dest.values[l*dest.stride+k] = mat.values[srcRowOffset+l]
					}
				}
			}
			return
		}

		// Recursive split into 4 equalish quadrants using integer division
		rowMid := (i + rowEnd) / 2
		colMid := (j + colEnd) / 2

		// round rowMid up to the next multiple of 8 to ensure alignment for SIMD processing
		if rowMid%8 != 0 {
			rowMid += 8 - (rowMid % 8)
			if rowMid > rowEnd {
				rowMid = rowEnd
			}
		}

		// round colMid up to the next multiple of 8 to ensure alignment for SIMD processing
		if colMid%8 != 0 {
			colMid += 8 - (colMid % 8)
			if colMid > colEnd {
				colMid = colEnd
			}
		}

		helper(i, j, rowMid, colMid)
		helper(i, colMid, rowMid, colEnd)
		helper(rowMid, j, rowEnd, colMid)
		helper(rowMid, colMid, rowEnd, colEnd)
	}

	helper(0, 0, mat.rows, mat.cols)
}

// transposeSIMD8x8 uses AVX2 256-bit vector registers and a butterfly shuffle
// network to transpose an 8x8 block in-register before writing out contiguously.
func transposeSIMD8x8(mat, dest *Matrix, i, j int) {
	// Scalar fallback if AVX2 is unavailable.
	if !archsimd.X86.AVX2() {
		for r := range 8 {
			for c := range 8 {
				dest.values[(j+c)*dest.stride+i+r] =
					mat.values[(i+r)*mat.stride+j+c]
			}
		}
		return
	}

	// Load eight rows and reinterpret their bits as uint32 vectors.
	a0 := archsimd.LoadFloat32x8(mat.values[(i+0)*mat.stride+j:]).ToBits()
	a1 := archsimd.LoadFloat32x8(mat.values[(i+1)*mat.stride+j:]).ToBits()
	a2 := archsimd.LoadFloat32x8(mat.values[(i+2)*mat.stride+j:]).ToBits()
	a3 := archsimd.LoadFloat32x8(mat.values[(i+3)*mat.stride+j:]).ToBits()
	a4 := archsimd.LoadFloat32x8(mat.values[(i+4)*mat.stride+j:]).ToBits()
	a5 := archsimd.LoadFloat32x8(mat.values[(i+5)*mat.stride+j:]).ToBits()
	a6 := archsimd.LoadFloat32x8(mat.values[(i+6)*mat.stride+j:]).ToBits()
	a7 := archsimd.LoadFloat32x8(mat.values[(i+7)*mat.stride+j:]).ToBits()

	// Stage 1: Interleave pairs of rows.
	t0 := a0.InterleaveLoGrouped(a1)
	t1 := a0.InterleaveHiGrouped(a1)
	t2 := a2.InterleaveLoGrouped(a3)
	t3 := a2.InterleaveHiGrouped(a3)
	t4 := a4.InterleaveLoGrouped(a5)
	t5 := a4.InterleaveHiGrouped(a5)
	t6 := a6.InterleaveLoGrouped(a7)
	t7 := a6.InterleaveHiGrouped(a7)

	// Stage 2: Combine groups of four rows.
	a0 = t0.ConcatPermuteScalarsGrouped(0, 1, 4, 5, t2)
	a1 = t0.ConcatPermuteScalarsGrouped(2, 3, 6, 7, t2)
	a2 = t1.ConcatPermuteScalarsGrouped(0, 1, 4, 5, t3)
	a3 = t1.ConcatPermuteScalarsGrouped(2, 3, 6, 7, t3)
	a4 = t4.ConcatPermuteScalarsGrouped(0, 1, 4, 5, t6)
	a5 = t4.ConcatPermuteScalarsGrouped(2, 3, 6, 7, t6)
	a6 = t5.ConcatPermuteScalarsGrouped(0, 1, 4, 5, t7)
	a7 = t5.ConcatPermuteScalarsGrouped(2, 3, 6, 7, t7)

	// Stage 3: Exchange 128-bit lanes.
	b0 := a0.ConcatPermute128Scalars(0, 2, a4)
	b1 := a1.ConcatPermute128Scalars(0, 2, a5)
	b2 := a2.ConcatPermute128Scalars(0, 2, a6)
	b3 := a3.ConcatPermute128Scalars(0, 2, a7)
	b4 := a0.ConcatPermute128Scalars(1, 3, a4)
	b5 := a1.ConcatPermute128Scalars(1, 3, a5)
	b6 := a2.ConcatPermute128Scalars(1, 3, a6)
	b7 := a3.ConcatPermute128Scalars(1, 3, a7)

	// Store the transposed rows.
	b0.BitsToFloat32().Store(dest.values[(j+0)*dest.stride+i:])
	b1.BitsToFloat32().Store(dest.values[(j+1)*dest.stride+i:])
	b2.BitsToFloat32().Store(dest.values[(j+2)*dest.stride+i:])
	b3.BitsToFloat32().Store(dest.values[(j+3)*dest.stride+i:])
	b4.BitsToFloat32().Store(dest.values[(j+4)*dest.stride+i:])
	b5.BitsToFloat32().Store(dest.values[(j+5)*dest.stride+i:])
	b6.BitsToFloat32().Store(dest.values[(j+6)*dest.stride+i:])
	b7.BitsToFloat32().Store(dest.values[(j+7)*dest.stride+i:])
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

// matMulIKJ multiples two matrices by applying SIMD to basic IKJ
func matMulIKJ(mat, other, dest *Matrix) {
	for i := range mat.rows {
		for k := range mat.cols {
			a := simd.BroadcastFloat32s(mat.values[i*mat.stride+k])

			j := 0
			for ; j+width <= other.cols; j += width {
				v1 := simd.LoadFloat32s(dest.values[i*dest.stride+j : i*dest.stride+j+width])
				v2 := simd.LoadFloat32s(other.values[k*other.stride+j : k*other.stride+j+width])
				v := a.MulAdd(v2, v1)
				v.Store(dest.values[i*dest.stride+j:])
			}

			if j < other.cols {
				rem := other.cols - j
				v1, _ := simd.LoadFloat32sPart(dest.values[i*dest.stride+j : i*dest.stride+j+rem])
				v2, _ := simd.LoadFloat32sPart(other.values[k*other.stride+j : k*other.stride+j+rem])
				v := a.MulAdd(v2, v1)
				v.StorePart(dest.values[i*dest.stride+j : i*dest.stride+j+rem])
			}
		}
	}
}

func matMulIJK(mat, other, dest *Matrix) {
	for i := range mat.rows {
		for j := 0; j < other.cols; j += width {
			acc := simd.BroadcastFloat32s(0)

			if j+width <= other.cols {
				for k := range mat.cols {
					a := simd.BroadcastFloat32s(mat.values[i*mat.stride+k])
					b := simd.LoadFloat32s(other.values[k*other.stride+j : k*other.stride+j+width])
					acc = a.MulAdd(b, acc)
				}
				acc.Store(dest.values[i*dest.stride+j : i*dest.stride+j+width])
			} else {
				for k := range mat.cols {
					a := simd.BroadcastFloat32s(mat.values[i*mat.stride+k])
					b, _ := simd.LoadFloat32sPart(other.values[k*other.stride+j : k*other.stride+other.cols])
					acc = a.MulAdd(b, acc)
				}
				acc.StorePart(dest.values[i*dest.stride+j : i*dest.stride+dest.cols])
			}
		}
	}
}

// matMulIJKUnrolled4 uses four independent SIMD accumulators per output row.
func matMulIJKUnrolled4(mat, other, dest *Matrix) {
	const vectorsPerBlock = 4
	blockWidth := vectorsPerBlock * width

	for i := range mat.rows {
		j := 0
		for ; j+blockWidth <= other.cols; j += blockWidth {
			acc0 := simd.BroadcastFloat32s(0)
			acc1 := simd.BroadcastFloat32s(0)
			acc2 := simd.BroadcastFloat32s(0)
			acc3 := simd.BroadcastFloat32s(0)

			for k := range mat.cols {
				a := simd.BroadcastFloat32s(mat.values[i*mat.stride+k])
				base := k*other.stride + j
				acc0 = a.MulAdd(simd.LoadFloat32s(other.values[base:base+width]), acc0)
				acc1 = a.MulAdd(simd.LoadFloat32s(other.values[base+width:base+2*width]), acc1)
				acc2 = a.MulAdd(simd.LoadFloat32s(other.values[base+2*width:base+3*width]), acc2)
				acc3 = a.MulAdd(simd.LoadFloat32s(other.values[base+3*width:base+4*width]), acc3)
			}

			base := i*dest.stride + j
			acc0.Store(dest.values[base : base+width])
			acc1.Store(dest.values[base+width : base+2*width])
			acc2.Store(dest.values[base+2*width : base+3*width])
			acc3.Store(dest.values[base+3*width : base+4*width])
		}

		for ; j < other.cols; j += width {
			end := min(j+width, other.cols)
			acc := simd.BroadcastFloat32s(0)

			for k := range mat.cols {
				a := simd.BroadcastFloat32s(mat.values[i*mat.stride+k])
				start := k*other.stride + j
				if end-j == width {
					acc = a.MulAdd(simd.LoadFloat32s(other.values[start:start+width]), acc)
				} else {
					b, _ := simd.LoadFloat32sPart(other.values[start : start+end-j])
					acc = a.MulAdd(b, acc)
				}
			}

			base := i*dest.stride + j
			if end-j == width {
				acc.Store(dest.values[base : base+width])
			} else {
				acc.StorePart(dest.values[base : base+end-j])
			}
		}
	}
}

// matMulIJKRows4 reuses each SIMD load from other across four output rows.
func matMulIJKRows4(mat, other, dest *Matrix) {
	const rowsPerBlock = 4

	for i := 0; i < mat.rows; i += rowsPerBlock {
		rowCount := min(rowsPerBlock, mat.rows-i)
		for j := 0; j < other.cols; j += width {
			end := min(j+width, other.cols)
			acc0 := simd.BroadcastFloat32s(0)
			acc1 := simd.BroadcastFloat32s(0)
			acc2 := simd.BroadcastFloat32s(0)
			acc3 := simd.BroadcastFloat32s(0)

			for k := range mat.cols {
				base := k*other.stride + j
				var b simd.Float32s
				if end-j == width {
					b = simd.LoadFloat32s(other.values[base : base+width])
				} else {
					b, _ = simd.LoadFloat32sPart(other.values[base : base+end-j])
				}

				a0 := simd.BroadcastFloat32s(mat.values[i*mat.stride+k])
				acc0 = a0.MulAdd(b, acc0)
				if rowCount > 1 {
					a1 := simd.BroadcastFloat32s(mat.values[(i+1)*mat.stride+k])
					acc1 = a1.MulAdd(b, acc1)
				}
				if rowCount > 2 {
					a2 := simd.BroadcastFloat32s(mat.values[(i+2)*mat.stride+k])
					acc2 = a2.MulAdd(b, acc2)
				}
				if rowCount > 3 {
					a3 := simd.BroadcastFloat32s(mat.values[(i+3)*mat.stride+k])
					acc3 = a3.MulAdd(b, acc3)
				}
			}

			for row := 0; row < rowCount; row++ {
				base := (i+row)*dest.stride + j
				var acc simd.Float32s
				switch row {
				case 0:
					acc = acc0
				case 1:
					acc = acc1
				case 2:
					acc = acc2
				case 3:
					acc = acc3
				}
				if end-j == width {
					acc.Store(dest.values[base : base+width])
				} else {
					acc.StorePart(dest.values[base : base+end-j])
				}
			}
		}
	}
}
