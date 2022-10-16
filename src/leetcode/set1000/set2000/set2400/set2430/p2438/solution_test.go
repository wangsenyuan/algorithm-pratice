package p2438

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := minimizeArrayValue(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 1}
	expect := 10
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 7, 1, 6}
	expect := 5
	runSample(t, A, expect)
}
