package shapes

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type ColoredCircle struct {
	Circle
	Color string
}
