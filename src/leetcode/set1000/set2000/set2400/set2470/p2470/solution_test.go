package p2470

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := subarrayLCM(A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 12, 9, 6}
	k := 3
	expect := 1
	runSample(t, nums, k, expect)
}
