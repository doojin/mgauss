package mgauss

// Matrix stores float values
type Matrix struct {
	vectors []Vector
}

// Adds vector to the slice of vecotrs
func (m *Matrix) AddVector(v Vector) {
	m.vectors = append(m.vectors, v)
}
