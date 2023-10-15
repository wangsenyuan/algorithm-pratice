package p2899

import "testing"

func runSample(t *testing.T, nums []int, l int, r int, expect int) {
	res := countSubMultisets(nums, l, r)

	if res != expect {
		t.Fatalf("sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 0, 1, 2, 3}
	l := 2
	r := 3
	expect := 9
	runSample(t, nums, l, r, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 2, 3}
	l := 1
	r := 9
	expect := 11
	runSample(t, nums, l, r, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2, 2, 3, 3}
	l := 9
	r := 10
	expect := 2
	runSample(t, nums, l, r, expect)
}
