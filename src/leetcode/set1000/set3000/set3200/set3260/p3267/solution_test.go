package p3267

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countPairs(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1023, 2310, 2130, 213}
	expect := 4
	runSample(t, a, expect)
}
