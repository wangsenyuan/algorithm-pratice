package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, r int, x []int, expect []float64) {
	res := solve(r, x)

	for i, y := range expect {
		y0 := res[i]
		dy := math.Abs(y - y0)
		eps := dy / math.Max(1, y)
		if eps > 1e-6 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	r := 2
	x := []int{5, 5, 6, 8, 3, 12}
	expect := []float64{2, 6.0, 9.87298334621, 13.3370849613, 12.5187346573, 13.3370849613}
	runSample(t, r, x, expect)
}
