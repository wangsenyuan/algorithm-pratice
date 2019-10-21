package p1230

import (
	"math"
	"testing"
)

func runSample(t *testing.T, prop []float64, target int, expect float64) {
	res := probabilityOfHeads(prop, target)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample %v %d, expect %f, but got %f", prop, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	prop := []float64{0.4}
	target := 1
	expect := 0.4
	runSample(t, prop, target, expect)
}

func TestSample2(t *testing.T) {
	prop := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	target := 0
	expect := 0.03125
	runSample(t, prop, target, expect)
}
