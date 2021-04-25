package p1837

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := sumBase(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 34, 6, 9)
}
