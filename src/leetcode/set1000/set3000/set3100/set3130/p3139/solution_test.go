package p3139

import "testing"

func runSample(t *testing.T, nums []int, cost1, cost2 int, expect int) {
	res := minCostToEqualizeArray(nums, cost1, cost2)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 1}
	cost1 := 5
	cost2 := 2
	expect := 15
	runSample(t, nums, cost1, cost2, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 14, 14, 15}
	cost1 := 2
	cost2 := 1
	expect := 20
	runSample(t, nums, cost1, cost2, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1000000}
	cost1 := 1000000
	cost2 := 1
	expect := 998993007
	runSample(t, nums, cost1, cost2, expect)
}
