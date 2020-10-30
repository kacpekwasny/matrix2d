package matrix2d

// Add adding two matricies
var Add = Pointer2Normal(add)

// Add adding two matricies
func add(m1 Matrix2D, m2 Matrix2D) (*Matrix2D, error) {
	if !AreSameSize(m1, m2) {
		return nil, errMatrixSizeError
	}
	l := [][]float64{}
	for y := 0; y < m1.lenY; y++ {
		l = append(l, addSlices(m1.M[y], m2.M[y]))
	}
	m, _ := InitMatrix2D(l)
	return m, nil
}

func addSlices(l1 []float64, l2 []float64) []float64 {
	// l1 and l2 should have same length
	l := make([]float64, len(l1), len(l1))
	for i := 0; i < len(l1); i++ {
		l[i] = l1[i] + l2[i]
	}
	return l
}
