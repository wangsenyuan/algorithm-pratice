package p1542

import "testing"

func runSample(t *testing.T, arr []int, target int, expect int) {
	res := maxNonOverlapping(arr, target)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1}
	target := 2
	expect := 2
	runSample(t, arr, target, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{-1, 3, 5, 1, 4, 2, -9}
	target := 6
	expect := 2
	runSample(t, arr, target, expect)
}
