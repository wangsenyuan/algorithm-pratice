package p964

import "testing"

func runSample(t *testing.T, x int, target int, expect int) {
	res := leastOpsExpressTarget(x, target)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", x, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 19, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 501, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 100000000, 3)
}
