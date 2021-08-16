package p1967

import "testing"

func runSample(t *testing.T, nums []int) {
	res := rearrangeArray(nums)

	if !check(res) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func check(nums []int) bool {
	for i := 1; i < len(nums)-1; i++ {
		a, b, c := nums[i-1], nums[i], nums[i+1]
		if 2*b == a+c {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 12, 10, 7, 6, 17, 14, 4, 13}
	runSample(t, nums)
}

func TestSample2(t *testing.T) {
	nums := []int{15, 7, 13, 6, 3, 11, 14, 1, 20}
	runSample(t, nums)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	runSample(t, nums)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 4, 0, 2, 5}
	runSample(t, nums)
}

func TestSample5(t *testing.T) {
	nums := []int{1, 2, 3}
	runSample(t, nums)
}
