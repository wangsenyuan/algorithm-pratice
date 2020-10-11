package p416

import "testing"

func runSample(t *testing.T, arr []int, expect bool) {
	res := canPartition(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 5, 11, 5}
	expect := true
	runSample(t, arr, expect)
}
