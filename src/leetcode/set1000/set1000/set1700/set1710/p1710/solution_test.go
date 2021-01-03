package p1710

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countPairs(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	deliciousness := []int{1, 3, 5, 7, 9}
	expect := 4
	runSample(t, deliciousness, expect)
}

func TestSample2(t *testing.T) {
	deliciousness := []int{1, 1, 1, 3, 3, 3, 7}
	expect := 15
	runSample(t, deliciousness, expect)
}

func TestSample3(t *testing.T) {
	deliciousness := []int{2, 14, 11, 5, 1744, 2352, 0, 1, 1300, 2796, 0, 4, 376, 1672, 73, 55, 2006, 42, 10, 6, 0, 2, 2, 0, 0, 1, 0, 1, 0, 2, 271, 241, 1, 63, 1117, 931, 3, 5, 378, 646, 2, 0, 2, 0, 15, 1}
	expect := 174
	runSample(t, deliciousness, expect)
}
