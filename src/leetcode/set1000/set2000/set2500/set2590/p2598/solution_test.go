package p2598

import "testing"

func runSample(t *testing.T, nums []int, value int, expect int) {
	res := findSmallestInteger(nums, value)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -10, 7, 13, 6, 8}
	value := 5
	expect := 4
	runSample(t, nums, value, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, -10, 7, 13, 6, 8}
	value := 7
	expect := 2
	runSample(t, nums, value, expect)
}
