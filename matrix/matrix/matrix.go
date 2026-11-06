package matrix

import "fmt"

// Matrix is a struct representing a 2D matrix of float32 values
type Matrix struct {
	values []float32
	rows   int
	cols   int
	stride [2]int
}

// NewMatrix returns an empty matrix with the given number of rows and columns
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		values: make([]float32, rows*cols),
		rows:   rows,
		cols:   cols,
		stride: [2]int{cols, 1},
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

// Get returns the value at the specified row and column
func (m *Matrix) Get(row, col int) float32 {
	return m.values[row*m.stride[0]+col*m.stride[1]]
}

// Set sets the value at the specified row and column
func (m *Matrix) Set(row, col int, value float32) {
	m.values[row*m.stride[0]+col*m.stride[1]] = value
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
		copy(values[i], m.values[i*m.stride[0]:(i+1)*m.stride[0]])
	}
	return values
}

// Add adds two matrices and returns a new matrix with the result
func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("matrix dimensions do not match: %dx%d vs %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	res := NewMatrix(m.rows, m.cols)
	add(m.values, other.values, res.values)

	return res
}
