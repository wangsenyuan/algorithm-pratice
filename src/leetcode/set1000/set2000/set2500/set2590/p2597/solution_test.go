package p2597

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := beautifulSubsets(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4, 6}
	k := 2
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{590, 136, 844, 976, 670, 485, 794, 114, 434, 82, 245, 673, 738, 416, 252, 1000, 518, 520, 1, 622}
	k := 999
	expect := 786431
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, 4, 5, 7, 2, 1}
	k := 2
	expect := 35
	runSample(t, nums, k, expect)
}
