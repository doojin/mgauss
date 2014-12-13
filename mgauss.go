package mgauss

// solveEquation calculates value of X
func solveEquation(xs map[int]float64, v Vector, dividend float64, index int) {
	coeff := v.Get(index)

	for i := index + 1; i < len(v.values); i++ {
		dividend -= xs[i] * v.Get(i)
	}

	xs[index] = dividend / coeff
}
