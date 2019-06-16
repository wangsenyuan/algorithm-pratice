package p823

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := numFactoredBinaryTrees(A)

	if expect != res {
		t.Errorf("sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 4}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 4, 5, 10}, 7)
}
