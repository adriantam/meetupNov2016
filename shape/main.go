package main

import (
	"fmt"
	"math"
)

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

// PaintCost takes in the interface Shape as parameter
// Thus, one does not need to write separate function for
// Square, Circle, etc.
func PaintCost(s Shape) int {
	return 500 * int(s.Area())
}

func ArrowProbability(distance int, s Shape) float64 {
	if distance != 0 {
		return math.Log10(s.Area()/float64(distance)) / 100
	}
	return 0
}

func main() {
	shapes := []Shape{Circle{radius: 9.3}, Square{side: 3.5}}
	for _, s := range shapes {
		fmt.Println("Cost of paint is ", PaintCost(s))
		fmt.Println("Chance of shot is  ", ArrowProbability(10, s))
	}
}
