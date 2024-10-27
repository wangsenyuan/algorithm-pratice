package p3337

import "testing"

func runSample(t *testing.T, s string, x int, nums []int, expect int) {
	res := lengthAfterTransformations(s, x, nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcyy"
	x := 2
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}
	expect := 7
	runSample(t, s, x, nums, expect)
}

func TestSample2(t *testing.T) {
	s := "azbk"
	x := 1
	nums := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	expect := 8
	runSample(t, s, x, nums, expect)
}
