package p2269

import "testing"

func runSample(t *testing.T, num int, k int, expect int) {
	res := divisorSubstrings(num, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 240, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 430043, 2, 2)
}
