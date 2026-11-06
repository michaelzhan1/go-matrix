package matrix

func naiveAdd(mat, other, dest []float32) {
	for i := 0; i < len(mat); i++ {
		dest[i] = mat[i] + other[i]
	}
}
