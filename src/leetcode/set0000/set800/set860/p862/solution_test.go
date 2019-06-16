package p862

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := shortestSubarray(A, K)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1}
	K := 1
	expect := 1
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2}
	K := 4
	expect := -1
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, -1, 2}
	K := 3
	expect := 3
	runSample(t, A, K, expect)
}

func TestSample4(t *testing.T) {
	A := []int{48, 99, 37, 4, -31}
	K := 140
	expect := 2
	runSample(t, A, K, expect)
}

func TestSample5(t *testing.T) {
	A := []int{-28, 81, -20, 28, -29}
	K := 89
	expect := 3
	runSample(t, A, K, expect)
}
