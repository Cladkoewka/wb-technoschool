package level1tasks

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	xDiff := math.Abs(p.x - other.x)
	yDiff := math.Abs(p.y - other.y)
	return math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
}

func Task24() {
	p1 := NewPoint(0, 4) 
	p2 := NewPoint(3, 0) 
	fmt.Println(p1.Distance(p2)) 
}