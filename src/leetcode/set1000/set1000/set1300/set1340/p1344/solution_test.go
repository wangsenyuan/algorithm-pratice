package p1344

import (
	"math"
	"testing"
)

func runSample(t *testing.T, hour, min int, expect float64) {
	res := angleClock(hour, min)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %d %d, expect %f, but got %f", hour, min, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 12, 30, 165)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 30, 75)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 15, 7.5)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 50, 155)
}
