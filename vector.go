package mgauss

// Vector stores float numbers and provides methods to manipulate them
type Vector struct {
	values []float64
}

// NewVector creates new vector
func NewVector(values []float64) (v Vector) {
	v.values = values
	return
}

// Get returns vector's element by it's index
func (v Vector) Get(index int) float64 {
	return v.values[index]
}
