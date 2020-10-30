package matrix2d

func multiplySlices(l1 []float64, l2 []float64) []float64 {
	// l1 and l2 should have same length
	for i := 0; i < len(l1); i++ {
		l1[i] = l1[i] * l2[i]
	}
	return l1
}

func multiplySlices2F(l1 []float64, l2 []float64) float64 {
	// l1 and l2 shoud have same length
	sum := 0.0
	for i := 0; i < len(l1); i++ {
		sum = sum + l1[i]*l2[i]
	}
	return sum
}

// Multiply two matrices
func Multiply(m1 Matrix2D, m2 Matrix2D) (*Matrix2D, error) {
	if !Multiplicable(m1, m2) {
		return nil, errMatrixSizeError
	}
	// 2DMatrix
	m, err := InitEmptyMatrix2D(m1.lenY, m2.lenX)
	if err != nil {
		return nil, err
	}
	xl1 := m1.SliceY()
	yl2 := m2.SliceX()

	for y, xl := range xl1 {
		for x, yl := range yl2 {
			m.M[y][x] = multiplySlices2F(xl, yl)
		}
	}
	return m, nil
}
