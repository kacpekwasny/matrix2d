package matrix2d

// SliceY returns matrix every horizontal slice
// normal
func SliceY(m Matrix2D) [][]float64 {
	ll := make([][]float64, len(m.M), len(m.M))
	copy(ll, m.M)
	return ll
}

// SliceX returns matrix every vertical slice
// the not normal
func SliceX(m Matrix2D) [][]float64 {
	l := [][]float64{}
	for i := 0; i < len(m.M[0]); i++ {
		// vertical slice
		l = append(l, []float64{})
		for k := 0; k < len(m.M); k++ {
			l[i] = append(l[i], m.M[k][i])
		}
	}
	return l
}
