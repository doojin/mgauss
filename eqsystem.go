package mgauss

// EqSystem represents a system of linear equations
type EqSystem struct {
	mA Matrix
	mB Matrix
}

// NewEqSystem creates new linear equation system
func NewEqSystem(mA Matrix, mB Matrix) (sys EqSystem) {
	sys.mA = mA
	sys.mB = mB
	return
}
