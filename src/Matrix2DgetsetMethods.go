package matrix2d

// GetXlen ok
func (m *Matrix2D) GetXlen() int {
	return m.lenX
}

// GetYlen ok
func (m *Matrix2D) GetYlen() int {
	return m.lenY
}

// GetCorrect correct atribute
func (m *Matrix2D) GetCorrect() bool {
	return m.correct
}
