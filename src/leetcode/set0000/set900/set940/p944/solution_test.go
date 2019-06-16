package p944

import "testing"

func runSample(t *testing.T, A []string, expect int) {
	res := minDeletionSize(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []string{"cba", "daf", "ghi"}
	runSample(t, A, 1)
}

func TestSample2(t *testing.T) {
	A := []string{"a", "b"}
	runSample(t, A, 0)
}

func TestSample3(t *testing.T) {
	A := []string{"zyx", "wvu", "tsr"}
	runSample(t, A, 3)
}
