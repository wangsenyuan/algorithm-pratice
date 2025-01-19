package p3428

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := minMaxSums(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2
	expect := 24
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 0, 6}
	k := 1
	expect := 22
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	expect := 12
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{10, 5, 9, 9, 10, 10, 7, 7, 9, 6, 9, 6, 7, 6, 4, 9, 8, 4, 2, 0, 0, 3, 9, 3, 10, 3, 1, 9, 8, 2, 8, 2, 0, 7, 7, 6, 4, 6, 7, 3, 2, 5, 6, 6, 5, 0, 5, 7, 8, 1}
	k := 29
	expect := 428790069
	runSample(t, nums, k, expect)
}
