package p1551

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := minOperations(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}
func TestSample2(t *testing.T) {
	runSample(t, 6, 9)
}
