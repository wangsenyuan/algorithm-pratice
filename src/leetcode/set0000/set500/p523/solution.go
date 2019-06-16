package main

import "fmt"

func main() {
	//nums := []int{23, 2, 6, 4, 7}
	//k := 6
	nums := []int{2, 1, 3}
	k := 2
	fmt.Println(checkSubarraySum(nums, k))
}

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	lx := make(map[int]int)
	lx[0] = -1
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]

		if k != 0 {
			sum = sum % k
		}

		if j, found := lx[sum]; found {
			if i-j > 1 {
				return true
			}
		} else {
			lx[sum] = i
		}

	}

	return false
}
