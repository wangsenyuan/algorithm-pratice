package p952

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := largestComponentSize(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 6, 15, 35}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{20, 50, 9, 63}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 3, 6, 7, 4, 12, 21, 39}
	expect := 8
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 3, 6, 7, 4, 12, 39}
	expect := 6
	runSample(t, A, expect)
}
