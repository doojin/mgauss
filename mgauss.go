package mgauss

// solveEquation calculates value of X
func solveEquation(xs map[int]float64, v Vector, dividend float64) {
	var index int
	var coeff float64

	for xIndex, xCoeff := range v.values {
		if xCoeff != 0 {
			coeff, index = xCoeff, xIndex
			break
		}
	}

	for i := index + 1; i < len(v.values); i++ {
		dividend -= xs[i] * v.Get(i)
	}

	xs[index] = dividend / coeff
}
