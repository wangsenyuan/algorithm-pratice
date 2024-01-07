package p2996

import "testing"

func runSample(t *testing.T, x int, y int, expect int) {
	res := minimumOperationsToMakeEqual(x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 26, 1, 3)
}
