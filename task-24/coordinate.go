package main

import (
	"fmt"
	"math"
)

// Point - x, y encapsulated coordinates
type Point struct {
	x float64
	y float64
}

// NewPoint - create a new point with x, y coordinates
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// String - Print coordinates to the console
func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f", p.x, p.y)
}

// Distance - calculate the distance between two coordinates
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(10, 10)
	p2 := NewPoint(25, 14)

	fmt.Println(p1)
	fmt.Println()
	fmt.Println(p2)
	fmt.Println("distance between two point is: ", Distance(p1, p2))
}
