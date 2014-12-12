package mgauss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEqSystem_ShouldCreateEquationSystemWithPassedMatrixes(t *testing.T) {
	mA := Matrix{}
	mA.AddVector(NewVector([]float64{1, 2, 3}))
	mA.AddVector(NewVector([]float64{-1, -2, -3}))
	mA.AddVector(NewVector([]float64{0.1, 0.2, 0.3}))
	mB := Matrix{}
	mB.AddVector(NewVector([]float64{4}))
	mB.AddVector(NewVector([]float64{-4}))
	mB.AddVector(NewVector([]float64{0.4}))

	eqSystem := NewEqSystem(mA, mB)

	assert.Equal(t, mA, eqSystem.mA, "Should be equal")
	assert.Equal(t, mB, eqSystem.mB, "Should be equal")
}
