package matrix

import (
	"fmt"
	"testing"
)

func benchmarkTransposeNoPadding(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrix(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrix(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transpose(mat, dest, rows, cols, rowTile, colTile)
	}
}

func benchmarkTranspose(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrixWithPadding(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrixWithPadding(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transpose(mat, dest, rows, cols, rowTile, colTile)
	}
}

func BenchmarkTransposeSquareTile(b *testing.B) {
	size := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	b.Run("no-padding", func(b *testing.B) {
		for _, s := range size {
			b.Run(fmt.Sprintf("%dx%d", s, s), func(b *testing.B) {
				benchmarkTransposeNoPadding(b, s, s)
			})
		}
	})

	b.Run("with-padding", func(b *testing.B) {
		for _, s := range size {
			b.Run(fmt.Sprintf("%dx%d", s, s), func(b *testing.B) {
				benchmarkTranspose(b, s, s)
			})
		}
	})
}

func BenchmarkTranspose(b *testing.B) {
	tiles := [][2]int{{16, 4}, {32, 4}, {64, 4}, {128, 4}, {16, 8}, {32, 8}, {64, 8}, {128, 8}}
	b.Run("no-padding", func(b *testing.B) {
		for _, tile := range tiles {
			b.Run(fmt.Sprintf("%dx%d", tile[0], tile[1]), func(b *testing.B) {
				benchmarkTransposeNoPadding(b, tile[0], tile[1])
			})
		}
	})

	b.Run("with-padding", func(b *testing.B) {
		for _, tile := range tiles {
			b.Run(fmt.Sprintf("%dx%d", tile[0], tile[1]), func(b *testing.B) {
				benchmarkTranspose(b, tile[0], tile[1])
			})
		}
	})
}
