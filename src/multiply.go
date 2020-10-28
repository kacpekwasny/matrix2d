package matrix2d

func multiplyLists(l1 []float64, l2 []float64) []float64 {
	// l1 and l2 should have same length
	l := len(l1)
	i := 0
	for i < l {
		l1[i] = l1[i] * l2[i]
		i++
	}
	return l1
}
