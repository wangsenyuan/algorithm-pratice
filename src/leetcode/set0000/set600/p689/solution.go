package main

import "fmt"

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)

	sums := make([]int, n-k+1)

	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	sums[0] = sum
	for i := k; i < n; i++ {
		sum += nums[i]
		sum -= nums[i-k]
		sums[i-k+1] = sum
	}

	left := make([]int, n-k+1)

	for i := 1; i < len(sums); i++ {
		if sums[i] > sums[left[i-1]] {
			left[i] = i
		} else {
			left[i] = left[i-1]
		}
	}

	right := make([]int, n-k+1)
	right[len(sums)-1] = len(sums) - 1
	for i := len(sums) - 2; i >= 0; i-- {
		if sums[i] >= sums[right[i+1]] {
			right[i] = i
		} else {
			right[i] = right[i+1]
		}
	}

	tmp := 0
	ans := -1
	for i := k; i+k < len(sums); i++ {
		a := left[i-k]
		b := right[i+k]
		if sums[a]+sums[i]+sums[b] > tmp {
			tmp = sums[a] + sums[i] + sums[b]
			ans = i
		}
	}

	res := []int{left[ans-k], ans, right[ans+k]}

	return res
}

func main() {
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
}
