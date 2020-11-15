package p5602

import "sort"

func minOperations(nums []int, x int) int {
	n := len(nums)
	left := make([]int, n)
	var best = -1

	for i := 0; i < n; i++ {
		left[i] = nums[i]
		if i > 0 {
			left[i] += left[i-1]
		}
		if left[i] == x {
			best = i + 1
		}
	}

	update := func(cnt int) {
		if best < 0 || best > cnt {
			best = cnt
		}
	}
	var sum int
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		if sum > x {
			break
		}
		if sum == x {
			update(n - i)
		} else {
			j := sort.SearchInts(left, x-sum)
			if j < i && left[j] == x-sum {
				update(j + 1 + (n - i))
			}
		}
	}

	return best
}
