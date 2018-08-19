package p891_

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := sumSubseqWidths(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 1, 3}, 6)
}
