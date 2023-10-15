package p2903

import "testing"

func runSample(t *testing.T, nums []int, indexDiff int, valueDiff int, expect bool) {
	res := findIndices(nums, indexDiff, valueDiff)

	if res[0] >= 0 != expect {
		t.Fatalf("Sample expect %t,b ut got %v", expect, res)
	}

	if !expect {
		return
	}
	if abs(res[1]-res[0]) < indexDiff || abs(nums[res[1]]-nums[res[0]]) < valueDiff {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 1, 4, 1}
	indexDiff := 2
	valueDiff := 4
	expect := true
	runSample(t, nums, indexDiff, valueDiff, expect)
}
