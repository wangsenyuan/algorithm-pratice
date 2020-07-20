package p1521

import "testing"

func runSample(t *testing.T, arr []int, target int, expect int) {
	res := closestToTarget(arr, target)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{9, 12, 3, 7, 15}
	target := 5
	expect := 2
	runSample(t, arr, target, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1000000, 1000000, 1000000}
	target := 1
	expect := 999999
	runSample(t, arr, target, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 4, 8, 16}
	target := 0
	expect := 0
	runSample(t, arr, target, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{70, 15, 21, 96}
	target := 4
	expect := 0
	runSample(t, arr, target, expect)
}
