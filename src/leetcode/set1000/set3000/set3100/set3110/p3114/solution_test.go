package p3114

import "testing"

func runSample(t *testing.T, nums []int, values []int, expect int) {
	res := minimumValueSum(nums, values)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 3, 3, 2}
	values := []int{0, 3, 3, 2}
	expect := 12
	runSample(t, nums, values, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 5, 7, 7, 7, 5}
	values := []int{0, 7, 5}
	expect := 17
	runSample(t, nums, values, expect)
}
func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	values := []int{2}
	expect := -1
	runSample(t, nums, values, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1}
	values := []int{1}
	expect := 1
	runSample(t, nums, values, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{1, 3, 2, 4, 7, 5, 3}
	values := []int{0, 5, 3}
	expect := 12
	runSample(t, nums, values, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{3, 1}
	values := []int{3, 1}
	expect := 4
	runSample(t, nums, values, expect)
}
