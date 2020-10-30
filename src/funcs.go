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
