package matrix2d

import "fmt"

// Fill matrix by using func that will be given:
// lenY, lenX, y, x, value
func (m *Matrix2D) Fill(f func(int, int, int, int, float64) float64) {
	for y := 0; y < m.lenY; y++ {
		for x := 0; x < m.lenX; x++ {
			m.M[y][x] = f(m.lenY, m.lenX, y, x, m.M[y][x])
		}
	}
}

// Print pretty
func (m *Matrix2D) Print() {
	for _, l := range m.M {
		fmt.Println(l)
	}
}
