package p5602

import "testing"

func runSample(t *testing.T, nums []int, x int, expect int) {
	res := minOperations(nums, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 4, 2, 3}
	x := 5
	expect := 2
	runSample(t, nums, x, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 6, 7, 8, 9}
	x := 4
	expect := -1
	runSample(t, nums, x, expect)
}
func TestSample3(t *testing.T) {
	nums := []int{3, 2, 20, 1, 1, 3}
	x := 10
	expect := 5
	runSample(t, nums, x, expect)
}
