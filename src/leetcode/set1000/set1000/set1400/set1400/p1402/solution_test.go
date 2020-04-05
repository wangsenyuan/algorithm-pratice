package p1402

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := maxSatisfaction(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{-1, -8, 0, 5, -9}, 14)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{4, 3, 2}, 20)
}
