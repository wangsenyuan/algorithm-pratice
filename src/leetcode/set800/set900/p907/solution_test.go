package p907

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := sumSubarrayMins(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 2, 4}
	expect := 17
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1}
	expect := 6
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{48, 87, 27}
	expect := 264
	runSample(t, A, expect)
}
