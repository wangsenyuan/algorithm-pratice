package p3082

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := sumOfPower(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 3
	expect := 6
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 3}
	k := 5
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 7
	expect := 0
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 6
	expect := 1
	runSample(t, nums, k, expect)
}
