package p5553

import "testing"

func runSample(t *testing.T, nums []int, quantity []int, expect bool) {
	res := canDistribute(nums, quantity)

	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	quantity := []int{2}
	expect := false
	runSample(t, nums, quantity, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	quantity := []int{2}
	expect := true
	runSample(t, nums, quantity, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 1, 2, 2}
	quantity := []int{2, 2}
	expect := true
	runSample(t, nums, quantity, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 1, 2, 3}
	quantity := []int{2, 2}
	expect := false
	runSample(t, nums, quantity, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	quantity := []int{2, 3}
	expect := true
	runSample(t, nums, quantity, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2}
	quantity := []int{4, 3, 4, 3, 3}
	expect := true
	runSample(t, nums, quantity, expect)
}
