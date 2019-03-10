package p1005

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := largestSumAfterKNegations(A, K)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 3}
	K := 1
	expect := 5
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, -1, 0, 2}
	K := 3
	expect := 6
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, -3, -1, 5, -4}
	K := 2
	expect := 13
	runSample(t, A, K, expect)
}

func TestSample4(t *testing.T) {
	A := []int{-2, 5, 0, 2, -2}
	K := 3
	expect := 11
	runSample(t, A, K, expect)
}
