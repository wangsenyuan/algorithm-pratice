package p1186

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := maximumSum(arr)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{-1, -1, -1}, -1)
}
