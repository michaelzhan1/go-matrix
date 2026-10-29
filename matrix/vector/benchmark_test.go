package vector

import "testing"

func BenchmarkVectorAddSmall(b *testing.B) {
	benchmarkVectorAdd(b, 1<<5)
}
func BenchmarkVectorAddMedium(b *testing.B) {
	benchmarkVectorAdd(b, 1<<10)
}

func BenchmarkVectorAddLarge(b *testing.B) {
	benchmarkVectorAdd(b, 1<<20)
}

func BenchmarkVectorAddNonMultiple(b *testing.B) {
	benchmarkVectorAdd(b, 999999)
}

func BenchmarkVectorAddSIMDSmall(b *testing.B) {
	benchmarkVectorAddSIMD(b, 1<<5)
}

func BenchmarkVectorAddSIMDMedium(b *testing.B) {
	benchmarkVectorAddSIMD(b, 1<<10)
}

func BenchmarkVectorAddSIMDLarge(b *testing.B) {
	benchmarkVectorAddSIMD(b, 1<<20)
}

func BenchmarkVectorAddSIMDNonMultiple(b *testing.B) {
	benchmarkVectorAddSIMD(b, 999999)
}

func benchmarkVectorAdd(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		for i := range v1 {
			res[i] = v1[i] + v2[i]
		}
	}
}

func benchmarkVectorAddSIMD(b *testing.B, size int) {
	v1 := NewRandomVector(size, 0, 100)
	v2 := NewRandomVector(size, 0, 100)
	res := make(Vector, size)

	b.ResetTimer()
	for b.Loop() {
		add(v1, v2, res)
	}
}
