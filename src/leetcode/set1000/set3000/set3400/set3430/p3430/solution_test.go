package p3430

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minMaxSubarraySum(nums, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2
	expect := 20
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, -3, 1}
	k := 2
	expect := -6
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-11, 17, -18, 0}
	k := 2
	// -11 + 17 + 17 + 17 + -18 + 0  = 22
	// -11 + -11 + 17 + -18 - 18 - -18 =
	expect := -37
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{11, 13, -20, -7, -12}
	k := 3
	expect := -100
	runSample(t, nums, k, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{-13, -16, -1, -19}
	k := 3
	expect := -201
	runSample(t, nums, k, expect)
}
