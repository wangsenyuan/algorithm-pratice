package p3336

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := subsequencePairCount(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	expect := 10
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 20, 30}
	expect := 2
	runSample(t, a, expect)
}
