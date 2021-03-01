package p1776

import (
	"math"
	"testing"
)

func runSample(t *testing.T, cars [][]int, expect []float64) {
	res := getCollisionTimes(cars)

	for i := 0; i < len(res); i++ {
		if math.Abs(res[i]-expect[i]) > 1e-6 {
			t.Fatalf("Sample expect %v, but got %v, diff at %d", expect, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	cars := [][]int{
		{1, 2}, {2, 1}, {4, 3}, {7, 2},
	}
	expect := []float64{
		1.00000, -1.00000, 3.00000, -1.00000,
	}
	runSample(t, cars, expect)
}

func TestSample2(t *testing.T) {
	cars := [][]int{
		{3, 4}, {5, 4}, {6, 3}, {9, 1},
	}
	expect := []float64{
		2.00000, 1.00000, 1.50000, -1.00000,
	}
	runSample(t, cars, expect)
}
