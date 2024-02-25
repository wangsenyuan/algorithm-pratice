package p3047

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := earliestSecondToMarkIndices(a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 2, 0}
	changeIndices := []int{2, 2, 2, 2, 3, 2, 2, 1}
	expect := 8
	runSample(t, nums, changeIndices, expect)
}
