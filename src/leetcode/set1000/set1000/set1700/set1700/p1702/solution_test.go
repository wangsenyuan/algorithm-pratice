package p1702

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minMoves(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 0, 0, 1, 0, 1}
	k := 2
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 0, 0, 0, 0, 1, 1}
	k := 3
	expect := 5
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 0, 1}
	k := 2
	expect := 0
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 1, 0, 1, 0, 0, 1}
	k := 4
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{1, 0, 0, 1, 0, 1}
	k := 2
	expect := 1
	runSample(t, nums, k, expect)
}
