package mgauss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveEquation_ShouldSolveEquationWithOneVariableCorrectly(t *testing.T) {
	xs := map[int]float64{}
	v := NewVector([]float64{0, 0, 0, 0, 1})
	divider := 5.0

	solveEquation(xs, v, divider, 4)

	assert.Equal(t, map[int]float64{4: 5.0}, xs)
}

func Test_solveEquation_ShouldSolveEquationWithTwoVariablesCorrectly(t *testing.T) {
	xs := map[int]float64{2: 5.0}
	v := NewVector([]float64{2, 0, 3, 0, 0})
	divider := -9.0

	solveEquation(xs, v, divider, 0)

	assert.Equal(t, map[int]float64{2: 5.0, 0: -12.0}, xs)
}
