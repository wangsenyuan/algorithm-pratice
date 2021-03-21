package p1802

import "testing"

func runSample(t *testing.T, n int, index int, maxSum int, expect int) {
	res := maxValue(n, index, maxSum)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	index := 2
	maxSum := 6
	expect := 2
	runSample(t, n, index, maxSum, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	index := 1
	maxSum := 10
	expect := 3
	runSample(t, n, index, maxSum, expect)
}
