package mgauss

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_NewVector_ShouldFillNewVectorWithCorrectValues(t *testing.T) {
	values := []float64{-1, 0.5, 1}

	v := NewVector(values)

	assert.Equal(t, []float64{-1, 0.5, 1}, v.values, "Must be equal")
}

func Test_NewVector_VectorValuesMustBeEmptyArrayIfEmptyArrayWasPassedAsArgument(t *testing.T) {
	v := NewVector([]float64{})

	assert.Equal(t, []float64{}, v.values, "Must be equal")
}
