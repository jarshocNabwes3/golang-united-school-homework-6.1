package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (triangle *Triangle) CalcArea() (area float64) {
	side := triangle.Side
	area = math.Sqrt(3) / 4 * side * side

	return
}

func (triangle *Triangle) CalcPerimeter() (perimeter float64) {
	side := triangle.Side
	perimeter = side * 3

	return
}
