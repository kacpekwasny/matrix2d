package matrix2d

// AreSameSize checks two matrices
func AreSameSize(m1 Matrix2D, m2 Matrix2D) bool {
	if m1.lenX == m2.lenX && m1.lenY == m2.lenY {
		return true
	}
	return false
}

// Multiplicable they can be multiplicated
func Multiplicable(m1 Matrix2D, m2 Matrix2D) bool {
	if m1.lenX == m2.lenY {
		return true
	}
	return false
}

// IsCorrect if every X row has same length
func IsCorrect(l [][]float64) bool {
	Xlength := len(l[0])
	for i := 0; i < len(l); i++ {
		if Xlength != len(l[i]) {
			return false
		}
	}
	return true
}
