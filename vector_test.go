package mgauss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVector_ShouldFillNewVectorWithCorrectValues(t *testing.T) {
	values := []float64{-1, 0.5, 1}

	v := NewVector(values)

	assert.Equal(t, []float64{-1, 0.5, 1}, v.values)
}

func Test_NewVector_VectorValuesMustBeEmptyArrayIfEmptyArrayWasPassedAsArgument(t *testing.T) {
	v := NewVector([]float64{})

	assert.Equal(t, []float64{}, v.values)
}

func Test_Get_ShouldReturnCorrectElementFromVecotr(t *testing.T) {
	v := Vector{}
	v.values = []float64{-0.5, 0, 0.5}

	assert.Equal(t, -0.5, v.Get(0))
	assert.Equal(t, 0, v.Get(1))
	assert.Equal(t, 0.5, v.Get(2))
}

func Test_Multiplicate_ShouldMultiplicateEachValueOfVectorCorrectly(t *testing.T) {
	v := NewVector([]float64{-0.5, 0, 0.5, 1})

	v.Multiplicate(-0.75)

	assert.Equal(t, []float64{0.375, 0, -0.375, -0.75}, v.values)
}

func Test_Substract_ShouldSubstractOneVectorFromAnotherCorrectly(t *testing.T) {
	v1 := NewVector([]float64{2, -1, 0, 0.5})
	v2 := NewVector([]float64{3, 2, -5, -2.5})

	v1.Subtract(v2)

	assert.Equal(t, []float64{-1, -3, 5, 3}, v1.values)
}

func Test_Length_ShouldReturnCorrectVectorLength(t *testing.T) {
	v1 := NewVector([]float64{1, 2, 3, 4, 5})

	length := v1.Length()

	assert.Equal(t, 5, length)
}
