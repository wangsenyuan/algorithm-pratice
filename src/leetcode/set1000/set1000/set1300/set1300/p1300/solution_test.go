package p1300

import "testing"

func runSample(t *testing.T, arr []int, target int, expect int) {
	res := findBestValue(arr, target)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{4, 9, 3}
	target := 10
	expect := 3
	runSample(t, arr, target, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 3, 5}
	target := 10
	expect := 5
	runSample(t, arr, target, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{60864, 25176, 27249, 21296, 20204}
	target := 56803
	expect := 11361
	runSample(t, arr, target, expect)
}
