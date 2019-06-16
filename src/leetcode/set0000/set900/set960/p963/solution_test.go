package p963

import (
	"math"
	"testing"
)

func runSample(t *testing.T, points [][]int, expect float64) {
	res := minAreaFreeRect(points)

	if math.Abs(res-expect) > 1e-5 {
		t.Errorf("Sample %v, expect %f, but got %f", points, expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{1, 2}, {2, 1}, {1, 0}, {0, 1}}
	expect := 2.0
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{0, 1}, {2, 1}, {1, 1}, {1, 0}, {2, 0}}
	expect := 1.0
	runSample(t, points, expect)
}
