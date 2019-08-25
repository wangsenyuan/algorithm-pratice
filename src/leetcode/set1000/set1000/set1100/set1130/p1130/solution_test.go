package p1130

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := mctFromLeafValues(arr)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{6, 2, 4}
	runSample(t, arr, 32)
}
