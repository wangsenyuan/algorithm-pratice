package p3473

import "testing"

func runSample(t *testing.T, nums []int, k int, m int, expect int) {
	res := maxSum(nums, k, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, -1, 3, 3, 4}
	k := 2
	m := 2
	expect := 13
	runSample(t, a, k, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-10, 3, -1, -2}
	k := 4
	m := 1
	expect := -10
	runSample(t, a, k, m, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-2, -10, 15, 12, 8, 11, 5}
	k := 3
	m := 2
	expect := 41
	runSample(t, a, k, m, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-8, 1, -8, 6, -9, 5}
	k := 1
	m := 3
	expect := 2
	runSample(t, a, k, m, expect)
}
