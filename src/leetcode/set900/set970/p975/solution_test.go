package p975

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := oddEvenJumps(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 13, 12, 14, 15}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 1, 1, 4}
	expect := 3
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 1, 3, 4, 2}
	expect := 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{19, 56, 23, 96}
	expect := 3
	runSample(t, A, expect)
}
