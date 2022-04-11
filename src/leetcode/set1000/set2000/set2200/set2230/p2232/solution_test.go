package p2232

import "testing"

func runSample(t *testing.T, k int, nums []int, expect int) {
	res := maximumProduct(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 4}
	k := 5
	expect := 20
	runSample(t, k, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{6, 3, 3, 2}
	k := 2
	expect := 216
	runSample(t, k, nums, expect)
}
