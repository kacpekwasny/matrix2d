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
