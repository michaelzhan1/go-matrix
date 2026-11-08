package matrix

import (
	"fmt"
	"testing"
)

func benchmarkTranspose(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrix(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrix(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transpose(mat, dest, rows, cols, rowTile, colTile)
	}
}

func benchmarkTransposeWithPadding(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrixWithPadding(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrixWithPadding(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transposePadded(mat, dest, rows, cols, rowTile, colTile)
	}
}

func BenchmarkTransposeSquareTile(b *testing.B) {
	size := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, s := range size {
		b.Run(fmt.Sprintf("tile-size-%dx%d", s, s), func(b *testing.B) {
			benchmarkTranspose(b, s, s)
		})
	}
}

func BenchmarkTranspose(b *testing.B) {
	tiles := [][2]int{{1, 1}, {2, 2}, {4, 4}, {8, 8}, {16, 8}, {16, 16}, {32, 8}, {32, 32}, {64, 8}, {64, 64}, {128, 8}}
	for _, tile := range tiles {
		b.Run(fmt.Sprintf("tile-size-%dx%d", tile[0], tile[1]), func(b *testing.B) {
			benchmarkTranspose(b, tile[0], tile[1])
		})
	}

	for _, tile := range tiles {
		b.Run(fmt.Sprintf("tile-size-%dx%d-with-padding", tile[0], tile[1]), func(b *testing.B) {
			benchmarkTransposeWithPadding(b, tile[0], tile[1])
		})
	}
}
