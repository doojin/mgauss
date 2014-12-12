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

	assert.Equal(t, mA, eqSystem.mA)
	assert.Equal(t, mB, eqSystem.mB)
}

func Test_swapRows_ShouldSwapRowsInBothMatrix(t *testing.T) {
	mA := Matrix{}
	mA.AddVector(NewVector([]float64{0, 0, 0, 0}))
	mA.AddVector(NewVector([]float64{1, 1, 1, 1}))
	mB := Matrix{}
	mB.AddVector(NewVector([]float64{2}))
	mB.AddVector(NewVector([]float64{3}))
	eqSystem := NewEqSystem(mA, mB)

	eqSystem.swapRows(0, 1)

	assert.Equal(t, []float64{1, 1, 1, 1}, eqSystem.mA.Row(0).values)
	assert.Equal(t, []float64{0, 0, 0, 0}, eqSystem.mA.Row(1).values)
	assert.Equal(t, []float64{3}, eqSystem.mB.Row(0).values)
	assert.Equal(t, []float64{2}, eqSystem.mB.Row(1).values)
}

func Test_itersCount_ShouldReturnCorrectIterationsCount(t *testing.T) {
	mA := Matrix{}
	mA.AddVector(NewVector([]float64{}))
	mA.AddVector(NewVector([]float64{}))
	mA.AddVector(NewVector([]float64{}))
	eqSystem := NewEqSystem(mA, Matrix{})

	itersCount := eqSystem.itersCount()

	assert.Equal(t, 2, itersCount)
}

func Test_doIteration_ShouldDoIterationCorrectly(t *testing.T) {
	mA := Matrix{}
	mA.AddVector(NewVector([]float64{2, -1, 1}))
	mA.AddVector(NewVector([]float64{8, 0, -2}))
	mA.AddVector(NewVector([]float64{1, 1, 1}))
	mB := Matrix{}
	mB.AddVector(NewVector([]float64{3}))
	mB.AddVector(NewVector([]float64{2}))
	mB.AddVector(NewVector([]float64{6}))
	eqSystem := NewEqSystem(mA, mB)

	eqSystem.doIteration(0)

	assert.Equal(t, []float64{1, 0, -0.25}, mA.Row(0).values)
	assert.Equal(t, []float64{0, -1, 1.5}, mA.Row(1).values)
	assert.Equal(t, []float64{0, 1, 1.25}, mA.Row(2).values)
	assert.Equal(t, []float64{0.25}, mB.Row(0).values)
	assert.Equal(t, []float64{2.5}, mB.Row(1).values)
	assert.Equal(t, []float64{5.75}, mB.Row(2).values)

	eqSystem.doIteration(1)

	assert.Equal(t, []float64{1, 0, -0.25}, mA.Row(0).values)
	assert.Equal(t, []float64{0, 1, -1.5}, mA.Row(1).values)
	assert.Equal(t, []float64{0, 0, 2.75}, mA.Row(2).values)
	assert.Equal(t, []float64{0.25}, mB.Row(0).values)
	assert.Equal(t, []float64{-2.5}, mB.Row(1).values)
	assert.Equal(t, []float64{8.25}, mB.Row(2).values)
}

func Test_transform_ShouldTransformCorrectlyIntoTriangularMatrix(t *testing.T) {
	mA := Matrix{}
	mA.AddVector(NewVector([]float64{0, 0, 3}))
	mA.AddVector(NewVector([]float64{0, 2, 0}))
	mA.AddVector(NewVector([]float64{1, 1, 1}))
	mB := Matrix{}
	mB.AddVector(NewVector([]float64{1}))
	mB.AddVector(NewVector([]float64{2}))
	mB.AddVector(NewVector([]float64{3}))
	eqSystem := NewEqSystem(mA, mB)

	eqSystem.transform()

	assert.Equal(t, []float64{1, 1, 1}, eqSystem.mA.Row(0).values)
	assert.Equal(t, []float64{0, 2, 0}, eqSystem.mA.Row(1).values)
	assert.Equal(t, []float64{0, 0, 3}, eqSystem.mA.Row(2).values)
	assert.Equal(t, []float64{3}, eqSystem.mB.Row(0).values)
	assert.Equal(t, []float64{2}, eqSystem.mB.Row(1).values)
	assert.Equal(t, []float64{1}, eqSystem.mB.Row(2).values)
}
