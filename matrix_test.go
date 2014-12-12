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
