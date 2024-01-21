package p3010

import "testing"

func runSample(t *testing.T, nums []int, k int, dist int, expect int) {
	res := minimumCost(nums, k, dist)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 2, 6, 4, 2}
	k := 3
	dist := 3
	expect := 5
	runSample(t, nums, k, dist, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{10, 1, 2, 2, 2, 1}
	k := 4
	dist := 3
	expect := 15
	runSample(t, nums, k, dist, expect)
}
func TestSample3(t *testing.T) {
	nums := []int{10, 8, 18, 9}
	k := 3
	dist := 1
	expect := 36
	runSample(t, nums, k, dist, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 6, 5, 8, 11, 10, 6}
	k := 5
	dist := 3
	expect := 31
	runSample(t, nums, k, dist, expect)
}
