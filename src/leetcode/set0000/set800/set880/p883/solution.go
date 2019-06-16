package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Solution struct {
	x, y, r float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{x_center, y_center, radius}
}

func (this *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64()) * this.r

	// degree
	d := rand.Float64() * 2 * math.Pi

	x := r * math.Cos(d)
	y := r * math.Sin(d)

	return []float64{x + this.x, y + this.y}
}

func main() {
	sln := Constructor(1, 0, 0)
	for i := 0; i < 1000; i++ {
		res := sln.RandPoint()
		x, y := res[0], res[1]
		fmt.Printf("%.7f %.7f\n", x, y)
	}
}
