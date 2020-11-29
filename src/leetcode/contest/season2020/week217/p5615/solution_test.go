package p5615

import "testing"

func runSample(t *testing.T, nums []int, limit int, expect int) {
	res := minMoves(nums, limit)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", nums, limit, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 4, 3}
	limit := 4
	expect := 1
	runSample(t, nums, limit, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 2, 1}
	limit := 2
	expect := 2
	runSample(t, nums, limit, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 1, 2}
	limit := 2
	expect := 0
	runSample(t, nums, limit, expect)
}
