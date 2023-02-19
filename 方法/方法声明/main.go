package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p1 := Point{1, 2}
	q1 := Point{4, 6}

	fmt.Println(Distance(p1, q1))

	fmt.Println(p1.Distance(q1))

	fmt.Println(q1.Distance(p1))

}
