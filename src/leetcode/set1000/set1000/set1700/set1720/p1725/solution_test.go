package p1725

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := tupleSameProduct(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 4, 6}
	expect := 8
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 4, 5, 10}
	expect := 16
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 3, 4, 6, 8, 12}
	expect := 40
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{2, 3, 5, 7}
	expect := 0
	runSample(t, nums, expect)
}
