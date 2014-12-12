package mgauss

import "math"

// Matrix stores float values
type Matrix struct {
	vectors []Vector
}

// AddVector appends vector to the slice of vecotrs
func (m *Matrix) AddVector(v Vector) {
	m.vectors = append(m.vectors, v)
}

// Col return matrix column
func (m *Matrix) Col(colIndex int) (v Vector) {
	for _, vector := range m.vectors {
		value := vector.Get(colIndex)
		v.values = append(v.values, value)
	}
	return
}

// mainElement returns the "main element" of matrix as well as position of
// vector it belongs to
func (m Matrix) mainElement(colIndex int) (v float64, pos int) {
	col := m.Col(colIndex)
	v, pos = math.Abs(col.Get(0)), 0
	for i, val := range col.values {
		if math.Abs(val) > v {
			v, pos = val, i
		}
	}
	return
}
