package mgauss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddVector_ShouldAppendVectorElementToTheVectorSlice(t *testing.T) {
	m := Matrix{}
	v := NewVector([]float64{-0.5, 0, 0.5})

	assert.Equal(t, m.vectors, []Vector(nil), "Should be equal")

	m.AddVector(v)

	assert.Equal(t, m.vectors, []Vector{v}, "Should be equal")
}

func Test_Col_ShouldReturnVectorOfMatrixColumnElements(t *testing.T) {
	m := Matrix{}
	m.vectors = []Vector{
		Vector{[]float64{-1, 2.5, -3, 4}},
		Vector{[]float64{5, -6, 7, -8.5}},
		Vector{[]float64{-9, 10.5, -11, 12}},
		Vector{[]float64{13, -14, 0, -16.5}},
	}

	assert.Equal(t, Vector{[]float64{-1, 5, -9, 13}}, m.Col(0), "Should be equal")
	assert.Equal(t, Vector{[]float64{2.5, -6, 10.5, -14}}, m.Col(1), "Should be equal")
	assert.Equal(t, Vector{[]float64{-3, 7, -11, 0}}, m.Col(2), "Should be equal")
	assert.Equal(t, Vector{[]float64{4, -8.5, 12, -16.5}}, m.Col(3), "Should be equal")
}
