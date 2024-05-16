package main

import "testing"

func runSample(t *testing.T, nums []int, expect string) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{6, 5, 2, 2, 5, 3}
	expect := "Alice"
	runSample(t, nums, expect)
}
