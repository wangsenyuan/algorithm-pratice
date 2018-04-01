package p808

import (
	"testing"
	"math"
)

func runSample(t *testing.T, N int, expect float64) {
	res := soupServings(N)
	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("sample %d, expect %f, but got %f", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 50, 0.625)
}
