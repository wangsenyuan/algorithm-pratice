package p3265

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countPairs(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 12, 8, 5, 5, 1, 20, 3, 10, 10, 5, 5, 5, 5, 1}
	expect := 27
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{30, 12, 3, 17, 21}
	expect := 2
	runSample(t, a, expect)
}
