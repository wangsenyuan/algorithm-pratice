package p2111

import "testing"

func runSample(t *testing.T, arr []int, k int, expect int) {
	res := kIncreasing(arr, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{4, 1, 5, 2, 6, 2}
	k := 2
	expect := 0
	runSample(t, arr, k, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	k := 1
	expect := 4
	runSample(t, arr, k, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{4, 1, 5, 2, 6, 2}
	k := 3
	expect := 2
	runSample(t, arr, k, expect)
}
