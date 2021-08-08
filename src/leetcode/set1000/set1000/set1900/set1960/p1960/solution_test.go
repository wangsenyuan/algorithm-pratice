package p1960

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minSpaceWastedKResizing(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{10, 20}
	k := 0
	expect := 10
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 20, 30}
	k := 1
	expect := 10
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{10, 20, 15, 30, 20}
	k := 2
	expect := 15
	runSample(t, nums, k, expect)
}
