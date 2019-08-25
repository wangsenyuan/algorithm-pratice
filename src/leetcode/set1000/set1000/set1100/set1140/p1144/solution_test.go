package p1144

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := movesToMakeZigzag(A)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{9, 6, 1, 6, 2}, 4)
}
