package p1478

import "testing"

func runSample(t *testing.T, arr []int, target int, expect int) {
	res := minSumOfLengths(arr, target)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{3, 2, 2, 4, 3}
	target := 3
	expect := 2
	runSample(t, arr, target, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2, 2, 3, 2, 6, 7, 2, 1, 4, 8}
	target := 5
	expect := 4
	runSample(t, arr, target, expect)
}
