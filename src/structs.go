package matrix2d

// Matrix2D f
type Matrix2D struct {
	//  [0,1,0,4,0],
	//  [0,0,1,5,1],
	//  [3,5,5,1,2],
	//  lenY = 3
	//  lenX = 5
	M       [][]float64
	correct bool
	lenY    int
	lenX    int
}
