package matrix

import (
	"fmt"
	"testing"
)

func benchmarkTransposeTileNoPadding(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrix(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrix(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for range b.N {
		transposeTile(mat, dest, rows, cols, rowTile, colTile)
	}
}

func benchmarkTransposeTile(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrixWithPadding(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrixWithPadding(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for range b.N {
		transposeTile(mat, dest, rows, cols, rowTile, colTile)
	}
}

func benchmarkTransposeRecursive(b *testing.B, rowTile, colTile int) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrixWithPadding(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrixWithPadding(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for range b.N {
		transposeRecursive(mat, dest, rows, cols, rowTile, colTile)
	}
}

func benchmarkTransposeSIMD(b *testing.B) {
	rows := 1024
	cols := 1024
	mat := NewRandomMatrixWithPadding(rows, cols, -1.0, 1.0)
	dest := NewRandomMatrixWithPadding(cols, rows, -1.0, 1.0)

	b.ResetTimer()
	for range b.N {
		transposeSIMD(mat, dest, rows, cols)
	}
}

func BenchmarkTransposeSquareTile(b *testing.B) {
	size := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	b.Run("no-padding", func(b *testing.B) {
		for _, s := range size {
			b.Run(fmt.Sprintf("%dx%d", s, s), func(b *testing.B) {
				benchmarkTransposeTileNoPadding(b, s, s)
			})
		}
	})

	b.Run("with-padding", func(b *testing.B) {
		for _, s := range size {
			b.Run(fmt.Sprintf("%dx%d", s, s), func(b *testing.B) {
				benchmarkTransposeTile(b, s, s)
			})
		}
	})
}

func BenchmarkTransposeRectangleTile(b *testing.B) {
	tiles := [][2]int{{16, 4}, {32, 4}, {64, 4}, {128, 4}, {16, 8}, {32, 8}, {64, 8}, {128, 8}}
	b.Run("no-padding", func(b *testing.B) {
		for _, tile := range tiles {
			b.Run(fmt.Sprintf("%dx%d", tile[0], tile[1]), func(b *testing.B) {
				benchmarkTransposeTileNoPadding(b, tile[0], tile[1])
			})
		}
	})

	b.Run("with-padding", func(b *testing.B) {
		for _, tile := range tiles {
			b.Run(fmt.Sprintf("%dx%d", tile[0], tile[1]), func(b *testing.B) {
				benchmarkTransposeTile(b, tile[0], tile[1])
			})
		}
	})
}

func BenchmarkTransposeMethods(b *testing.B) {
	tiles := [][2]int{{64, 8}, {128, 8}, {16, 16}, {32, 32}, {64, 64}, {128, 128}}
	b.Run("transposeTile", func(b *testing.B) {
		for _, tile := range tiles {
			b.Run(fmt.Sprintf("%dx%d", tile[0], tile[1]), func(b *testing.B) {
				benchmarkTransposeTile(b, tile[0], tile[1])
			})
		}
	})

	b.Run("transposeRecursive", func(b *testing.B) {
		for _, tile := range tiles {
			b.Run(fmt.Sprintf("%dx%d", tile[0], tile[1]), func(b *testing.B) {
				benchmarkTransposeRecursive(b, tile[0], tile[1])
			})
		}
	})

	b.Run("transposeSIMD", func(b *testing.B) {
		benchmarkTransposeSIMD(b)
	})
}
