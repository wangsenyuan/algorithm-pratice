package p1095

import "testing"

func runSample(t *testing.T, arr []int, target int, expect int) {
	mountain := &MountainArray{arr: arr}

	res := findInMountainArray(target, mountain)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 3, 1}
	target := 3
	expect := 2
	runSample(t, arr, target, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{0, 1, 2, 4, 2, 1}
	target := 3
	expect := -1
	runSample(t, arr, target, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 5, 2}
	target := 0
	expect := -1
	runSample(t, arr, target, expect)
}
