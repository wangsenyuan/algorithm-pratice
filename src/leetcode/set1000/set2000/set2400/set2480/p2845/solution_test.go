package p2845

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := pivotInteger(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 6)
}
