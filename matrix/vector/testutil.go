package vector

import "math/rand/v2"

// NewRandomVector returns a Vector of length n with random values between min and max.
func NewRandomVector(n int, min, max float32) Vector {
	v := make(Vector, n)
	for i := range v {
		v[i] = min + rand.Float32()*(max-min)
	}
	return v
}
