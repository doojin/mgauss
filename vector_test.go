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

func Test_Get_ShouldReturnCorrectElementFromVecotr(t *testing.T) {
	v := Vector{}
	v.values = []float64{-0.5, 0, 0.5}

	assert.Equal(t, -0.5, v.Get(0), "Should be equal")
	assert.Equal(t, 0, v.Get(1), "Should be equal")
	assert.Equal(t, 0.5, v.Get(2), "Should be equal")
}
