package p753

import "testing"

func runSample(t *testing.T, n int, k int, minL int) {
	res := crackSafe(n, k)
	if len(res) != minL {
		t.Errorf("the sample %d %d should give a string of length of %d, bug got %s of length %d", n, k, minL, res, len(res))
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, 5)
}
