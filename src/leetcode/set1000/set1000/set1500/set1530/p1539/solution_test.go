package p1539

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int) {
	res := findKthPositive(arr, k)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", arr, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	expect := 9
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	k := 2
	expect := 6
	runSample(t, arr, k, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{2, 3, 4}
	k := 1
	expect := 1
	runSample(t, arr, k, expect)
}
