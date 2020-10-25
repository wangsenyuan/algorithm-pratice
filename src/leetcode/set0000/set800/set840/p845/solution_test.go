package p845

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := longestMountain(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 1, 4, 7, 3, 2, 5}
	expect := 5
	runSample(t, arr, expect)
}
