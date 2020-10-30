package matrix2d

// Pointer2Normal takes a func that required func(*m1, *m2)
// returns func that requires func(m1, m2)
func Pointer2Normal(f func(Matrix2D, Matrix2D) (*Matrix2D, error)) func(*Matrix2D, *Matrix2D) (*Matrix2D, error) {
	return func(m1 *Matrix2D, m2 *Matrix2D) (*Matrix2D, error) {
		return f(*m1, *m2)
	}
}
