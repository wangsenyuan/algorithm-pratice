package p2452

import "testing"

func runSample(t *testing.T, nums []int, space int, expect int) {
	res := destroyTargets(nums, space)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 7, 8, 1, 1, 5}
	space := 2
	expect := 1
	runSample(t, nums, space, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 5, 3, 2, 2}
	space := 100
	expect := 2
	runSample(t, nums, space, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{625879766, 235326233, 250224393, 501422042, 683823101, 948619719, 680305710, 733191937, 182186779, 353350082}
	space := 4
	expect := 235326233
	runSample(t, nums, space, expect)
}
