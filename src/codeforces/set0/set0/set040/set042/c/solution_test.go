package main

import (
	"slices"
	"strconv"
	"testing"
)

func runSample(t *testing.T, nums []int) {
	arr := make([]int, len(nums))
	copy(arr, nums)
	res := solve(arr)

	if len(res) > 1000 {
		t.Fatalf("Sample result took too much operations, %d", len(res))
	}

	for _, op := range res {
		i, _ := strconv.Atoi(op[1:])
		i--
		if op[0] == '+' {
			nums[i]++
			nums[(i+1)%4]++
		} else {
			if nums[i]%2 != 0 || nums[(i+1)%4]%2 != 0 {
				t.Fatalf("result is wrong, it has div on odd number %v", nums)
			}
			nums[i] /= 2
			nums[(i+1)%4] /= 2
		}
	}
	if slices.Max(nums) != 1 {
		t.Fatalf("result is wrong, it has number %v", nums)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	runSample(t, nums)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 1, 2}
	runSample(t, nums)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 1, 2, 1}
	runSample(t, nums)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 4, 2}
	runSample(t, nums)
}

func TestSample5(t *testing.T) {
	nums := []int{3, 3, 1, 1}
	runSample(t, nums)
}

func TestSample6(t *testing.T) {
	nums := []int{2, 1, 2, 4}
	runSample(t, nums)
}

func TestSample8(t *testing.T) {
	nums := []int{307791796, 300139641, 49227316, 460665315}
	runSample(t, nums)
}

func TestSample9(t *testing.T) {
	nums := []int{2, 2, 16, 1}
	runSample(t, nums)
}

func TestSample10(t *testing.T) {
	nums := []int{420679048, 116660561, 519054676, 79846989}
	runSample(t, nums)
}
