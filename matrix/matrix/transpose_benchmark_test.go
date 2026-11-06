package matrix

import (
	"fmt"
	"testing"
)

func benchmarkTranspose(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := make([]float32, rows*cols)
	dest := make([]float32, rows*cols)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transpose(mat, dest, rows, cols, rowTile, colTile)
	}
}

func BenchmarkTranspose(b *testing.B) {
	tiles := [][2]int{{1, 1}, {2, 2}, {4, 4}, {8, 8}, {16, 8}, {16, 16}, {32, 8}, {32, 32}, {64, 8}, {64, 64}, {128, 8}}
	for _, tile := range tiles {
		b.Run(fmt.Sprintf("tile-size-%dx%d", tile[0], tile[1]), func(b *testing.B) {
			benchmarkTranspose(b, tile[0], tile[1])
		})
	}
}
