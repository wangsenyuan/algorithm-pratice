package p2556

import "testing"

func runSample(t *testing.T, A []int, lower, upper int, expect int64) {
	res := countFairPairs(A, lower, upper)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 1, 7, 4, 4, 5}
	lower := 3
	upper := 6
	var expect int64 = 6
	runSample(t, A, lower, upper, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 7, 9, 2, 5}
	lower := 11
	upper := 11
	var expect int64 = 1
	runSample(t, A, lower, upper, expect)
}
