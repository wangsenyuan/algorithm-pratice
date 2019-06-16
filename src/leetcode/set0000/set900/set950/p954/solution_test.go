package p954

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	res := canReorderDoubled(A)
	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 3, 6}
	runSample(t, A, false)
}

func TestSample2(t *testing.T) {
	A := []int{2, 1, 2, 6}
	runSample(t, A, false)
}

func TestSample3(t *testing.T) {
	A := []int{4, -2, 2, -4}
	runSample(t, A, true)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 4, 16, 8, 4}
	runSample(t, A, false)
}
