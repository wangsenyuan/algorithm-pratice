package p930

import "testing"

func runSample(t *testing.T, A []int, S int, expect int) {
	res := numSubarraysWithSum(A, S)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 0, 1, 0, 1}
	S := 2
	expect := 4
	runSample(t, A, S, expect)
}
