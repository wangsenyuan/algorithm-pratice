package p888_

import "testing"

func runSample(t *testing.T, A, B []int) {
	res := fairCandySwap(A, B)
	sa := sumArray(A)
	sb := sumArray(B)
	if sa-res[0]+res[1] != sb-res[1]+res[0] {
		t.Errorf("Sample %v %v, ans %v, is not correct", A, B, res)
	} else {
		t.Logf("Sample %v %v, got one ans %v", A, B, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 1}, []int{2, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2}, []int{2, 3})
}

func TestSample3(t *testing.T) {
	runSample(t, []int{2}, []int{1, 3})
}

func TestSample4(t *testing.T) {
	runSample(t, []int{1, 2, 5}, []int{2, 4})
}
