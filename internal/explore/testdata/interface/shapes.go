package shapes

// Shape is implemented by every concrete shape below. A call through this
// interface can't be resolved to one function statically — every
// implementation is a possible target.
type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return 3.14159 * c.Radius * c.Radius }

type Square struct{ Side float64 }

func (s Square) Area() float64 { return s.Side * s.Side }

func totalArea(shapes []Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}