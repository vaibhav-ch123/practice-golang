package square

type square struct {
	side float64
}

func New(side float64) square {
	s := square{side}
	return s
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Perimeter() float64 {
	return 4 * s.side
}
