package p795

import "testing"

func runSample(t *testing.T, A []int, L, R int, expect int) {
	res := numSubarrayBoundedMax(A, L, R)
	if res != expect {
		t.Errorf("sample %v %d %d, expect %d, but got %d", A, L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 4, 3}
	L := 2
	R := 3
	expect := 3
	runSample(t, A, L, R, expect)
}

func TestSample2(t *testing.T) {
	A := []int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}
	L := 32
	R := 69
	expect := 22
	runSample(t, A, L, R, expect)
}
