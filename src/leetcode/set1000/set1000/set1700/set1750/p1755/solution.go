package p1755

import "sort"

func minAbsDifference(nums []int, goal int) int {
	n := len(nums)
	// n <= 40
	h := n / 2
	first := nums[:h]
	second := nums[h:]
	arr1 := getSums(first)
	arr2 := getSums(second)
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	var best = abs(int64(goal))

	for _, cur1 := range arr1 {
		// cur1 + cur2 - goal
		// cur2   goal - cur1
		tmp := int64(goal) - cur1
		i := sort.Search(len(arr2), func(i int) bool {
			return arr2[i] >= tmp
		})
		if i < len(arr2) {
			best = min(best, abs(cur1+arr2[i]-int64(goal)))
		}
		// arr2[i-1] < tmp
		if i-1 >= 0 {
			best = min(best, abs(cur1+arr2[i-1]-int64(goal)))
		}

	}

	return int(best)
}

func getSums(arr []int) []int64 {
	n := len(arr)
	N := 1 << n
	res := make([]int64, 0, N)
	for mask := 0; mask < N; mask++ {
		var sum int64
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				sum += int64(arr[i])
			}
		}
		res = append(res, sum)
	}
	return res
}

func sum(arr []int) int64 {
	var res int64
	for _, num := range arr {
		res += int64(num)
	}
	return res
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
