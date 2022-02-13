package p2170

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := minimumOperations(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 3, 2, 4, 3}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 2, 2, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 3, 3, 4, 3, 3, 3}
	expect := 2
	runSample(t, nums, expect)
}
