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

// transform transforms matrix A into triangular matrix
func (sys EqSystem) transform() {
	for i := 0; i < len(sys.mA.vectors); i++ {
		sys.doIteration(i)
	}
}

func (sys EqSystem) doIteration(iter int) {
	_, vectorIndex := sys.mA.mainElement(iter)
	sys.swapRows(vectorIndex, iter)

	for i := iter + 1; i < len(sys.mA.vectors); i++ {
		currentRow := sys.mA.Row(i)

		if currentRow.Get(iter) == 0 {
			continue
		}

		coeff := currentRow.Get(iter) / sys.mA.Row(iter).Get(iter)

		sys.mA.Row(iter).Multiplicate(coeff)
		sys.mB.Row(iter).Multiplicate(coeff)

		sys.mA.Row(i).Subtract(sys.mA.Row(iter))
		sys.mB.Row(i).Subtract(sys.mB.Row(iter))
	}

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
