package p2475

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := unequalTriplets(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 4, 2, 4, 3}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 2, 2, 3, 3, 4}
	expect := 20
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{4, 4, 2, 4, 3}
	expect := 3
	runSample(t, A, expect)
}
