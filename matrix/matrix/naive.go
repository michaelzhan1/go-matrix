package matrix

func naiveAdd(mat, other, dest []float32) {
	for i := range mat {
		dest[i] = mat[i] + other[i]
	}
}

func naiveSub(mat, other, dest []float32) {
	for i := range mat {
		dest[i] = mat[i] - other[i]
	}
}

func naiveElemMul(mat, other, dest []float32) {
	for i := range mat {
		dest[i] = mat[i] * other[i]
	}
}

func naiveElemDiv(mat, other, dest []float32) {
	for i := range mat {
		dest[i] = mat[i] / other[i]
	}
}
