package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rectangle *Rectangle) CalcArea() (area float64) {
	Height := rectangle.Height
	Weight := rectangle.Weight
	area = Height * Weight

	return
}

func (rectangle *Rectangle) CalcPerimeter() (perimeter float64) {
	Height := rectangle.Height
	Weight := rectangle.Weight

	perimeter = (Weight + Height) * 2

	return
}
