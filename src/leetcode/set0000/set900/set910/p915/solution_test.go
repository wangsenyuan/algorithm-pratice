package p915

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := partitionDisjoint(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{5, 0, 3, 8, 6}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 1, 0, 6, 12}, 4)
}
