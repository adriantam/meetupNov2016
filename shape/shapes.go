package main

type Shape interface {
	Area() float64
}
type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
