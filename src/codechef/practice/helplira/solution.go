package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var coords [6]int

	fmt.Scanf("%d", &n)
	var x, y int
	var mx, my float64

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d %d %d", &coords[0], &coords[1], &coords[2], &coords[3], &coords[4], &coords[5])
		area := calculateArea(coords)

		if i == 0 || area <= mx {
			x = i
			mx = area
		}

		if i == 0 || area >= my {
			y = i
			my = area
		}
	}

	fmt.Printf("%d %d\n", x+1, y+1)
}
func calculateArea(coords [6]int) float64 {
	a := calculateSideLength(coords[0], coords[1], coords[2], coords[3])
	b := calculateSideLength(coords[0], coords[1], coords[4], coords[5])
	c := calculateSideLength(coords[2], coords[3], coords[4], coords[5])

	s := 0.5 * (a + b + c)

	areaSqar := s * (s - a) * (s - b) * (s - c)

	return areaSqar
}
func calculateSideLength(a int, b int, c int, d int) float64 {
	x := c - a
	y := d - b
	return math.Sqrt(float64(x*x + y*y))
}
