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

// Length returns vector's length
func (v Vector) Length() int {
	return len(v.values)
}

// Multiplicate multiplicates each value of vector and number
func (v Vector) Multiplicate(number float64) {
	for i := range v.values {
		v.values[i] *= number
	}
}

// Subtract subtracts one vector from another
func (v Vector) Subtract(vector Vector) {
	for i := range v.values {
		v.values[i] -= vector.values[i]
	}
}
