package p812

import (
	"math"
	"testing"
)

func runSample(t *testing.T, points [][]int, expect float64) {
	res := largestTriangleArea(points)
	if math.Abs(res-expect) > 0.0000001 {
		t.Errorf("sample %v, expect %f, but got %f", points, expect, res)
	}
}

func TestSample(t *testing.T) {
	points := [][]int{
		{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0},
	}

	runSample(t, points, 2)
}
