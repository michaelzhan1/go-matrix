package vector

// naiveAdd is a naive implementation of vector addition.
func naiveAdd(v1, v2, dest Vector) {
	for i := range v1 {
		dest[i] = v1[i] + v2[i]
	}
}

// naiveSub is a naive implementation of vector subtraction.
func naiveSub(v1, v2, dest Vector) {
	for i := range v1 {
		dest[i] = v1[i] - v2[i]
	}
}

// naiveMul is a naive implementation of vector multiplication.
func naiveMul(v1, v2, dest Vector) {
	for i := range v1 {
		dest[i] = v1[i] * v2[i]
	}
}

// naiveDiv is a naive implementation of vector division.
func naiveDiv(v1, v2, dest Vector) {
	for i := range v1 {
		dest[i] = v1[i] / v2[i]
	}
}

// naiveScalarAdd is a naive implementation of scalar addition.
func naiveScalarAdd(v Vector, scalar float32, dest Vector) {
	for i := range v {
		dest[i] = v[i] + scalar
	}
}

// naiveScalarSub is a naive implementation of scalar subtraction.
func naiveScalarSub(v Vector, scalar float32, dest Vector) {
	for i := range v {
		dest[i] = v[i] - scalar
	}
}

// naiveScalarMul is a naive implementation of scalar multiplication.
func naiveScalarMul(v Vector, scalar float32, dest Vector) {
	for i := range v {
		dest[i] = v[i] * scalar
	}
}

// naiveScalarDiv is a naive implementation of scalar division.
func naiveScalarDiv(v Vector, scalar float32, dest Vector) {
	for i := range v {
		dest[i] = v[i] / scalar
	}
}

// naiveDot is a naive implementation of the dot product.
func naiveDot(v1, v2 Vector) float32 {
	var sum float32
	for i := range v1 {
		sum += v1[i] * v2[i]
	}
	return sum
}

// naiveSum is a naive implementation of the sum of vector elements.
func naiveSum(v Vector) float32 {
	var sum float32
	for i := range v {
		sum += v[i]
	}
	return sum
}
