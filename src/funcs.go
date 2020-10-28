package matrix2d

/* RangeAll from Top Left to Right Bottom horizontally
func RangeAllTlRbH(m1 Matrix2D, m2 Matrix2D) (chan [2]float64, error) {
	if !AreSameSize(m1, m2) {
		return nil, errMatrixSizeError
	}
	c1 := make(chan [2]float64)
	i := 0
	l := len(m1.l)

	go func() {
		defer close(c1)
		for i<l {
			c1 <- []float64{}
		i++
		}
	}
	return
}
*/
