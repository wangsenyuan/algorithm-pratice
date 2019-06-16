package p961

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := repeatedNTimes(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3, 3}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 1, 2, 5, 3, 2}, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{5, 1, 5, 2, 5, 3, 5, 4}, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{9, 5, 6, 9}, 9)
}
