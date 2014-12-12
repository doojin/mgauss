package mgauss

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
