package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (circle *Circle) CalcArea() (area float64) {
	radius := circle.Radius
	area = math.Pi * radius * radius

	return
}

func (circle *Circle) CalcPerimeter() (perimeter float64) {
	radius := circle.Radius
	perimeter = math.Pi * radius * 2

	return
}
