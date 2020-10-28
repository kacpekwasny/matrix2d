package matrix2d

// Add adding two matricies
func Add(m1 Matrix2D, m2 Matrix2D) (*Matrix2D, error) {
	if !AreSameSize(m1, m2) {
		return nil, errMatrixSizeError
	}
	l := [][]float64{}
	y := 0
	for y < m1.lenY {
		l = append(l, addLists(m1.l[y], m2.l[y]))
		y++
	}
	m, _ := InitMatrix2D(l)
	return m, nil
}

func addLists(l1 []float64, l2 []float64) []float64 {
	// l1 and l2 should have same length
	l := len(l1)
	i := 0
	for i < l {
		l1[i] = l1[i] + l2[i]
		i++
	}
	return l1
}
