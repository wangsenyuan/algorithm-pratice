package p1004

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := longestOnes(A, K)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	expect := 6
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	K := 3
	expect := 10
	runSample(t, A, K, expect)
}
