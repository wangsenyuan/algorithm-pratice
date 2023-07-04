package main

import "testing"

func runSample(t *testing.T, nums []int, expect bool) {
	res := solve(nums)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}

	a := make(map[int]int)
	b := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if res[i] == 'A' {
			a[nums[i]] ^= 1
			if a[nums[i]] == 0 {
				delete(a, nums[i])
			}
		} else {
			b[nums[i]] ^= 1
			if b[nums[i]] == 0 {
				delete(b, nums[i])
			}
		}
	}

	if len(a) != len(b) {
		t.Fatalf("Sample result %v, not correct, with diff %d", res, len(a)-len(b))
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 5, 6, 6, 6, 7, 9, 11}
	expect := true
	runSample(t, nums, expect)
}
