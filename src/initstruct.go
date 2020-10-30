package matrix2d

// InitMatrix2D create from [][]float64
func InitMatrix2D(l [][]float64) (*Matrix2D, error) {
	m := Matrix2D{}
	m.M = l
	m.lenY = len(l)
	m.lenX = len(l[0])
	m.correct = IsCorrect(m.M)

	return &m, nil
}

// InitEmptyMatrix2D by specifying size
func InitEmptyMatrix2D(sizeY int, sizeX int) (*Matrix2D, error) {
	if sizeY < 0 || sizeX < 0 {
		return nil, errMatrixSizeError
	}
	m := Matrix2D{}
	m.lenY = sizeY
	m.lenX = sizeX

	m.M = make([][]float64, m.lenY, m.lenY)
	for i := 0; i < m.lenY; i++ {
		m.M[i] = make([]float64, m.lenX, m.lenX)
	}
	m.correct = IsCorrect(m.M)

	return &m, nil
}
