package p2208

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := halveArray(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 19, 8, 1}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 8, 20}
	expect := 3
	runSample(t, nums, expect)
}
