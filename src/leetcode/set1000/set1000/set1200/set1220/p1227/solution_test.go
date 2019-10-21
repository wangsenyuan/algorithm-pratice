package p1227

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, expect float64) {
	res := nthPersonGetsNthSeat(n)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d, expect %f, but got %f", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 0.5)
}
