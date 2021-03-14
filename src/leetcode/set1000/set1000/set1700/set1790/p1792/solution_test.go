package p1792

import (
	"math"
	"testing"
)

func runSample(t *testing.T, classes [][]int, extra int, expect float64) {
	res := maxAverageRatio(classes, extra)

	if math.Abs(res-expect) > 1e-5 {
		t.Errorf("Sample expect %f, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	classes := [][]int{
		{1, 2}, {3, 5}, {2, 2},
	}
	extra := 2
	expect := 0.78333
	runSample(t, classes, extra, expect)
}

func TestSample2(t *testing.T) {
	classes := [][]int{
		{2, 4}, {3, 9}, {4, 5}, {2, 10},
	}
	extra := 4
	expect := 0.53485
	runSample(t, classes, extra, expect)
}
