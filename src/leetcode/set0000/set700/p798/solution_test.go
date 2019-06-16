package p798

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := bestRotation(A)
	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 3, 1, 4, 0}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 0, 2, 4}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 1, 4, 0, 0}
	expect := 3
	runSample(t, A, expect)
}
