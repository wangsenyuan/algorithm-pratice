package p1394

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := numTeams(arr)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %v", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 5, 3, 4, 1}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 1, 3}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expect := 4
	runSample(t, arr, expect)
}
