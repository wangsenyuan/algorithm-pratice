package p1719

import "testing"

func runSample(t *testing.T, nums []int) {
	input := make([]int, len(nums))
	copy(input, nums)
	res := missingTwo(nums)

	n := len(nums) + 2
	seen := make([]bool, n)

	for _, num := range input {
		seen[num-1] = true
	}

	for _, num := range res {
		if seen[num-1] {
			t.Fatalf("Sample %v result %v, not correct", num, res)
		}
		seen[num-1] = true
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1}
	runSample(t, nums)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3}
	runSample(t, nums)
}
