package matrix2d

// Map returns new matrix with every element passed through func
func Map(m Matrix2D, f func(float64) float64) *Matrix2D {
	// mr - matrix for return
	mr, _ := InitEmptyMatrix2D(m.lenY, m.lenX)
	for y := 0; y < mr.lenY; y++ {
		for x := 0; x < mr.lenX; x++ {
			mr.M[y][x] = f(m.M[y][x])
		}
	}
	return mr
}

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
