package p801

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := minSwap(A, B)
	if res != expect {
		t.Errorf("sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 5, 4}
	B := []int{1, 2, 3, 7}
	expect := 1
	runSample(t, A, B, expect)
}
