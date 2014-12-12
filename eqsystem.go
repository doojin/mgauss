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

// swapRows swaps rows in both: matrix A and matrix B
func (sys EqSystem) swapRows(pos1 int, pos2 int) {
	sys.mA.SwapRows(pos1, pos2)
	sys.mB.SwapRows(pos1, pos2)
}

// itersCount returns iterations count needed to resolve an equation system
func (sys EqSystem) itersCount() int {
	return len(sys.mA.vectors) - 1
}
