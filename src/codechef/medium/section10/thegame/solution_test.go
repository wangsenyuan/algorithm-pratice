package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, R, C int, G []string, expect float64) {
	res := solve(R, C, G)

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample %d %d %v, expect %f,  but got %f", R, C, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C := 3, 3
	G := []string{
		"o#o",
		"oo#",
		"#oo",
	}
	expect := 1.166666667
	runSample(t, R, C, G, expect)
}
