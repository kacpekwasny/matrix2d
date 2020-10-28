package matrix2d

// InitMatrix2D create
func InitMatrix2D(l [][]float64) (*Matrix2D, error) {
	m := Matrix2D{}
	m.l = l
	m.lenY = len(l)
	m.lenX = len(l[0])
	m.correct = IsCorrect(m.l)

	return &m, nil
}

// IsCorrect if every X row has same length
func IsCorrect(l [][]float64) bool {
	i := 0
	ln := len(l)
	Xlength := len(l[0])
	for i < ln {
		if Xlength != len(l[i]) {
			return false
		}
		i++
	}
	return true
}

// RangeY ranges over every horizontal list
func (m *Matrix2D) RangeY() chan []float64 {
	c1 := make(chan []float64)
	go func() {
		defer close(c1)
		l := len(m.l)
		i := 0
		for i < l {
			c1 <- m.l[i]
			i++
		}
	}()
	return c1
}

// RangeX ranges over every vertical list
func (m *Matrix2D) RangeX() chan []float64 {
	c1 := make(chan []float64)
	go func() {
		defer close(c1)
		ln := len(m.l[0])
		i := 0
		for i < ln {
			// vertical list
			vls := []float64{}
			j := len(m.l)
			k := 0
			for k < j {
				vls = append(vls, m.l[j][i])
				j++
			}
			c1 <- vls
			i++
		}
	}()
	return c1
}
