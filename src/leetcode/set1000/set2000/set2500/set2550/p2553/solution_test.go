package p2553

import "testing"

func runSample(t *testing.T, weights []int, k int, expect int) {
	res := int(putMarbles(weights, k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	ws := []int{1, 3, 5, 1}
	k := 2
	expect := 4
	runSample(t, ws, k, expect)
}

func TestSample2(t *testing.T) {
	ws := []int{1, 3}
	k := 2
	expect := 0
	runSample(t, ws, k, expect)
}
