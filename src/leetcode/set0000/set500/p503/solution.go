package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)

	ans := make([]int, n)
	stack := make([]int, n)
	p := 0
	for i := 0; i < 2*n; i++ {
		for p > 0 && nums[stack[p-1]] < nums[i%n] {
			ans[stack[p-1]] = nums[i%n]
			p -= 1
		}
		if i < n {
			stack[p] = i
			p++
			ans[i % n] = -1
		}
	}

	return ans
}

/*

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	mx, mi := 0, -1

	for i, num := range nums {
		if mi == -1 || num > mx {
			mi = i
			mx = num
		}
	}

	n := len(nums)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nums[(mi+i+1)%n]
	}
	gt := make([]int, n)
	gt[n-1] = n

	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && xs[j] <= xs[i] {
			j = gt[j]
		}
		gt[i] = j
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		j := gt[i]
		k := (i + 1 + mi) % n
		if j == n {
			ans[k] = -1
		} else {
			ans[k] = xs[j]
		}
	}

	return ans
}
*/
