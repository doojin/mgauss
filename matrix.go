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

// Col return matrix column by it's index
func (m *Matrix) Col(index int) (v Vector) {
	for _, vector := range m.vectors {
		value := vector.Get(index)
		v.values = append(v.values, value)
	}
	return
}

// Row returns matrix row by it's index
func (m *Matrix) Row(index int) (v Vector) {
	return m.vectors[index]
}

// Copy returns a copy of matrix
func (m Matrix) Copy() Matrix {
	copy := Matrix{}
	copy.vectors = []Vector{}
	for _, v := range m.vectors {
		vectorValues := []float64{}
		for _, vectorValue := range v.values {
			vectorValues = append(vectorValues, vectorValue)
		}
		copy.vectors = append(copy.vectors, NewVector(vectorValues))
	}
	return copy
}

// SwapRows swaps two rows
func (m *Matrix) SwapRows(index1 int, index2 int) {
	m.vectors[index1], m.vectors[index2] = m.Row(index2), m.Row(index1)
}

// mainElement returns the "main element" of matrix as well as position of
// vector it belongs to
func (m Matrix) mainElement(index int) (v float64, pos int) {
	col := m.Col(index)
	v, pos = math.Abs(col.Get(index)), index
	for i := index; i < len(col.values); i++ {
		val := col.Get(i)
		if math.Abs(val) > v {
			v, pos = val, i
		}
	}
	return
}
