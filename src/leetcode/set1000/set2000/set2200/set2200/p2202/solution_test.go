package p2202

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maximumTop(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{99, 95, 68, 24, 18}
	k := 69
	expect := 99
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{18}
	k := 3
	expect := -1
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{73, 63, 62, 16, 95, 92, 93, 52, 89, 36, 75, 79, 67, 60, 42, 93, 93, 74, 94, 73, 35, 86, 96}
	k := 59
	expect := 96
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{91, 98, 17, 79, 15, 55, 47, 86, 4, 5, 17, 79, 68, 60, 60, 31, 72, 85, 25, 77, 8, 78, 40, 96, 76, 69, 95, 2, 42, 87, 48, 72, 45, 25, 40, 60, 21, 91, 32, 79, 2, 87, 80, 97, 82, 94, 69, 43, 18, 19, 21, 36, 44, 81, 99}
	k := 2
	expect := 91
	runSample(t, nums, k, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{94, 23, 58, 92, 3, 63, 68, 43, 8, 97, 54, 11, 63, 55, 73, 38, 22, 80, 45, 43, 25, 34, 67, 76, 80, 9, 30, 27, 50, 7, 57, 63, 63, 27, 46, 1, 50, 55, 99, 92, 73, 9, 57, 25, 59, 54, 100, 56, 64, 94, 99}
	k := 79
	expect := 100
	runSample(t, nums, k, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{1, 2, 3, 4, 1}
	k := 1000000000
	expect := 4
	runSample(t, nums, k, expect)
}

func TestSample7(t *testing.T) {
	nums := []int{0, 1, 2}
	k := 3
	expect := 1
	runSample(t, nums, k, expect)
}

func TestSample9(t *testing.T) {
	nums := []int{0, 1, 2}
	k := 5
	expect := 2
	runSample(t, nums, k, expect)
}
