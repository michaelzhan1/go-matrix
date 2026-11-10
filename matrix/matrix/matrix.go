package matrix

import "fmt"

// Matrix is a struct representing a 2D matrix of float32 values
type Matrix struct {
	values []float32
	rows   int
	cols   int
	stride int // account for padding
}

// NewMatrix returns an empty matrix with the given number of rows and columns
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		values: make([]float32, rows*cols),
		rows:   rows,
		cols:   cols,
		stride: cols,
	}
}

// NewMatrixWithPadding returns an empty matrix with the given number of rows and columns and padding.
func NewMatrixWithPadding(rows, cols int) *Matrix {
	pad := 0
	if cols%64 == 0 {
		pad = 8
	}

	stride := cols + pad
	return &Matrix{
		values: make([]float32, rows*stride),
		rows:   rows,
		cols:   cols,
		stride: stride,
	}
}

// From2DSlice creates a new Matrix from a 2D slice of float32 values
func From2DSlice(values [][]float32) *Matrix {
	rows := len(values)
	cols := len(values[0])

	m := NewMatrix(rows, cols)
	for i, row := range values {
		for j, val := range row {
			m.Set(i, j, val)
		}
	}

	return m
}

func From2DSliceWithPadding(values [][]float32) *Matrix {
	rows := len(values)
	cols := len(values[0])

	m := NewMatrixWithPadding(rows, cols)
	for i, row := range values {
		for j, val := range row {
			m.Set(i, j, val)
		}
	}

	return m
}

// Get returns the value at the specified row and column
func (m *Matrix) Get(row, col int) float32 {
	return m.values[row*m.stride+col]
}

// Set sets the value at the specified row and column
func (m *Matrix) Set(row, col int, value float32) {
	m.values[row*m.stride+col] = value
}

// Rows returns the number of rows in the matrix
func (m *Matrix) Rows() int {
	return m.rows
}

// Cols returns the number of columns in the matrix
func (m *Matrix) Cols() int {
	return m.cols
}

// Values returns the underlying 2D slice of float32 values
func (m *Matrix) Values() [][]float32 {
	values := make([][]float32, m.rows)
	for i := range m.rows {
		values[i] = make([]float32, m.cols)
		copy(values[i], m.values[i*m.stride:(i+1)*m.stride])
	}
	return values
}

// TransposeTile transposes the matrix in place using tiling
func (m *Matrix) TransposeTile() *Matrix {
	res := NewMatrix(m.cols, m.rows)
	transposeTile(m, res, m.rows, m.cols, 128, 8)
	return res
}

// TransposeRecursive transposes the matrix in place using recursive logic
func (m *Matrix) TransposeRecursive() *Matrix {
	res := NewMatrix(m.cols, m.rows)
	transposeRecursive(m, res, m.rows, m.cols, 16, 16)
	return res
}

// TransposeSIMD transposes the matrix in place using SIMD instructions
func (m *Matrix) TransposeSIMD() *Matrix {
	res := NewMatrix(m.cols, m.rows)
	transposeSIMD(m, res)
	return res
}

// Add adds two matrices and returns a new matrix with the result
func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("matrix dimensions do not match: %dx%d vs %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	res := NewMatrix(m.rows, m.cols)
	add(m, other, res)
	return res
}

// Sub subtracts two matrices and returns a new matrix with the result
func (m *Matrix) Sub(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("matrix dimensions do not match: %dx%d vs %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	res := NewMatrix(m.rows, m.cols)
	sub(m, other, res)
	return res
}

// ElemMul multiplies two matrices element-wise and returns a new matrix with the result
func (m *Matrix) ElemMul(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("matrix dimensions do not match: %dx%d vs %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	res := NewMatrix(m.rows, m.cols)
	elemMul(m, other, res)
	return res
}

// ElemDiv divides two matrices element-wise and returns a new matrix with the result
func (m *Matrix) ElemDiv(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("matrix dimensions do not match: %dx%d vs %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	res := NewMatrix(m.rows, m.cols)
	elemDiv(m, other, res)
	return res
}
