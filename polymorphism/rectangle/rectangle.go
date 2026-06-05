package rectangle

type rectangle struct {
	length  float64
	breadth float64
}

func New(l float64, b float64) rectangle {
	r := rectangle{l, b}
	return r
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r rectangle) Perimeter() float64 {
	return (2 * r.length) + (2 * r.breadth)
}
