package matrix

func naiveAdd(mat, other, dest *Matrix) {
	for i := range mat.values {
		dest.values[i] = mat.values[i] + other.values[i]
	}
}

func naiveSub(mat, other, dest *Matrix) {
	for i := range mat.values {
		dest.values[i] = mat.values[i] - other.values[i]
	}
}

func naiveElemMul(mat, other, dest *Matrix) {
	for i := range mat.values {
		dest.values[i] = mat.values[i] * other.values[i]
	}
}

func naiveElemDiv(mat, other, dest *Matrix) {
	for i := range mat.values {
		dest.values[i] = mat.values[i] / other.values[i]
	}
}

// naiveMatMul uses ikj ordering, where k is the shared dimension
func naiveMatMul(mat, other, dest *Matrix) {
	for i := range mat.rows {
		for k := range mat.cols {
			a := mat.values[i*mat.stride+k]
			for j := range other.cols {
				dest.values[i*dest.stride+j] += a * other.values[k*other.stride+j]
			}
		}
	}
}
