package p837

import (
	"testing"
	"math"
)

func runSample(t *testing.T, n int, K int, W int, expect float64) {
	res := new21Game(n, K, W)
	if math.Abs(res-expect) > 1e-5 {
		t.Errorf("Sample %d %d %d, expect %f, but got %f", n, K, W, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 21, 17, 10, 0.73278)
}
