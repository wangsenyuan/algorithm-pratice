package p1246

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := minimumMoves(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2}
	expect := 2
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 3, 4, 1, 5}
	expect := 3
	runSample(t, arr, expect)
}
