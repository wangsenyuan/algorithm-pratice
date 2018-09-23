package p908

import (
	"testing"
)

func runSample(t *testing.T, A []int, K int, expect int) {
	res := smallestRangeI(A, K)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1}
	K := 0
	expect := 0
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 10}
	K := 2
	expect := 6
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 6}
	K := 3
	expect := 0
	runSample(t, A, K, expect)
}
