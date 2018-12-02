package p949

import "testing"

func runSample(t *testing.T, A []int, expect string) {
	res := largestTimeFromDigits(A)
	if res != expect {
		t.Errorf("Sample %v, expect %s, but got %s", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3, 4}, "23:41")
}

func TestSample2(t *testing.T) {
	runSample(t, []int{5, 5, 5, 5}, "")
}
