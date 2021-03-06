package mgauss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddVector_ShouldAppendVectorElementToTheVectorSlice(t *testing.T) {
	m := Matrix{}
	v := NewVector([]float64{-0.5, 0, 0.5})

	assert.Equal(t, m.vectors, []Vector(nil))

	m.AddVector(v)

	assert.Equal(t, m.vectors, []Vector{v})
}

func Test_Col_ShouldReturnVectorOfMatrixColumnElements(t *testing.T) {
	m := Matrix{}
	m.vectors = []Vector{
		Vector{[]float64{-1, 2.5, -3, 4}},
		Vector{[]float64{5, -6, 7, -8.5}},
		Vector{[]float64{-9, 10.5, -11, 12}},
		Vector{[]float64{13, -14, 0, -16.5}},
	}

	assert.Equal(t, Vector{[]float64{-1, 5, -9, 13}}, m.Col(0))
	assert.Equal(t, Vector{[]float64{2.5, -6, 10.5, -14}}, m.Col(1))
	assert.Equal(t, Vector{[]float64{-3, 7, -11, 0}}, m.Col(2))
	assert.Equal(t, Vector{[]float64{4, -8.5, 12, -16.5}}, m.Col(3))
}

func Test_Row_ShouldReturnVecotrOfMatrixRowElements(t *testing.T) {
	m := Matrix{}
	m.vectors = []Vector{
		Vector{[]float64{-1, 2.5, -3, 4}},
		Vector{[]float64{5, -6, 7, -8.5}},
		Vector{[]float64{-9, 10.5, -11, 12}},
		Vector{[]float64{13, -14, 0, -16.5}},
	}

	assert.Equal(t, Vector{[]float64{-1, 2.5, -3, 4}}, m.Row(0))
	assert.Equal(t, Vector{[]float64{5, -6, 7, -8.5}}, m.Row(1))
	assert.Equal(t, Vector{[]float64{-9, 10.5, -11, 12}}, m.Row(2))
	assert.Equal(t, Vector{[]float64{13, -14, 0, -16.5}}, m.Row(3))
}

func Test_mainElement_ShouldReturnTheMainElementOfMatrixDependingOnColumnIndex(t *testing.T) {
	m := Matrix{}
	m.AddVector(NewVector([]float64{2, 3, 0.3}))
	m.AddVector(NewVector([]float64{-4, 2, 8}))
	m.AddVector(NewVector([]float64{-5.5, 0, 1}))

	val, index := m.mainElement(0)
	assert.Equal(t, -5.5, val)
	assert.Equal(t, 2, index)

	val, index = m.mainElement(1)
	assert.Equal(t, 2, val)
	assert.Equal(t, 1, index)

	val, index = m.mainElement(2)
	assert.Equal(t, 1, val)
	assert.Equal(t, 2, index)
}

func Test_mainElement_ShouldNotReturnElementFromZeroRowOnFirstIteration(t *testing.T) {
	m := Matrix{}
	m.AddVector(NewVector([]float64{9, 9, 9, 9}))
	m.AddVector(NewVector([]float64{0, 1, 1, 1}))
	m.AddVector(NewVector([]float64{0, 2, 2, 2}))
	m.AddVector(NewVector([]float64{0, -3, 3, 3}))

	val, index := m.mainElement(1)

	assert.Equal(t, val, -3)
	assert.Equal(t, index, 3)
}

func Test_SwapRows_ShouldSwapTwoRows(t *testing.T) {
	m := Matrix{}
	m.AddVector(NewVector([]float64{1, 1, 1}))
	m.AddVector(NewVector([]float64{2, 2, 2}))

	m.SwapRows(0, 1)

	assert.Equal(t, []float64{2, 2, 2}, m.Row(0).values)
	assert.Equal(t, []float64{1, 1, 1}, m.Row(1).values)
}

func Test_Copy_ShouldReturnACopyOfMatrix(t *testing.T) {
	m := Matrix{}
	m.AddVector(NewVector([]float64{1, 2, 3}))
	m.AddVector(NewVector([]float64{-0.5, 1.5, 0}))

	copy := m.Copy()

	assert.Equal(t, []float64{1, 2, 3}, copy.Row(0).values)
	assert.Equal(t, []float64{-0.5, 1.5, 0}, copy.Row(1).values)
}

func Test_Copy_WhenChangingCopiedMatrixOriginalMatrixShouldNotBeAffected(t *testing.T) {
	m := Matrix{}
	m.AddVector(NewVector([]float64{1, 2, 3}))
	m.AddVector(NewVector([]float64{-0.5, 1.5, 0}))

	copy := m.Copy()
	copy.Row(0).values[0] = 15
	copy.AddVector(NewVector([]float64{5, 5, 5}))

	assert.Equal(t, 2, len(m.vectors))
	assert.Equal(t, 3, len(copy.vectors))
	assert.Equal(t, []float64{1, 2, 3}, m.Row(0).values)
	assert.Equal(t, []float64{15, 2, 3}, copy.Row(0).values)
}
