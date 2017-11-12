package main

import "fmt"

func main() {
	findNums := []int{4, 1, 2}
	nums := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(findNums, nums))
}

func nextGreaterElement(findNums []int, nums []int) []int {
	if len(findNums) == 0 {
		return nil
	}
	n := len(nums)
	idx := make(map[int]int)

	stack := make([]int, n)
	p := 0
	gt := make([]int, n)

	for i, num := range nums {
		idx[num] = i
		gt[i] = n
		for p > 0 && num > nums[stack[p-1]] {
			gt[stack[p-1]] = i
			p--
		}
		stack[p] = i
		p++
	}

	ans := make([]int, len(findNums))

	for i := 0; i < len(findNums); i++ {
		j := gt[idx[findNums[i]]]
		if j == n {
			ans[i] = -1
		} else {
			ans[i] = nums[j]
		}
	}
	return ans
}

/*

func nextGreaterElement(findNums []int, nums []int) []int {
	if len(findNums) == 0 {
		return nil
	}
	n := len(nums)
	gt := make([]int, n)
	idx := make(map[int]int)
	gt[n-1] = n
	idx[nums[n-1]] = n - 1
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && nums[j] < nums[i] {
			j = gt[j]
		}
		gt[i] = j
		idx[nums[i]] = i
	}

	ans := make([]int, len(findNums))

	for i := 0; i < len(findNums); i++ {
		num := findNums[i]
		k := idx[num]
		j := gt[k]
		if j == n {
			ans[i] = -1
		} else {
			ans[i] = nums[j]
		}
	}
	return ans
}
*/
