package p2966

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minimumCost(nums)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 6
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 12, 13, 14, 15}
	expect := 11
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{22, 33, 22, 33, 22}
	expect := 22
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1321, 7284, 9346, 9460, 7099, 2796, 5887, 9351, 2278, 7590, 7627, 1552, 5864, 7409, 9356, 8480, 2765, 8036, 8473, 5573}
	expect := 42709
	runSample(t, a, expect)
}
