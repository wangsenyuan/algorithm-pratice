package p3095

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := sumOfPowers(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	k := 3
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 3, -1}
	k := 2
	expect := 10
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2}
	k := 2
	expect := 0
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{-16, -6, -12, -2}
	k := 4
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{1, 2, 3, 4, 6, 1, 9, 100, -1, -3, 1011}
	k := 3
	expect := 1356
	runSample(t, nums, k, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 6}
	k := 4
	expect := 16
	runSample(t, nums, k, expect)
}
