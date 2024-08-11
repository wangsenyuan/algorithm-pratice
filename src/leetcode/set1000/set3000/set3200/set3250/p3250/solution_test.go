package p3250

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countOfPairs(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 3, 2}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 5, 5, 5}
	expect := 126
	runSample(t, nums, expect)
}
