package p5559

func minimumMountainRemovals(nums []int) int {
	n := len(nums)

	lis := make([]int, n)
	stack := make([]int, n)
	var p int

	// LIS for i
	for i := 0; i < n; i++ {
		j := search(p, func(j int) bool {
			return stack[j] >= nums[i]
		})
		lis[i] = j
		if j == p {
			stack[p] = nums[i]
			p++
		} else {
			stack[j] = nums[i]
		}
	}
	var best = n
	p = 0
	for i := n - 1; i >= 0; i-- {
		j := search(p, func(j int) bool {
			return stack[j] >= nums[i]
		})

		if j > 0 && lis[i] > 0 {
			tmp := n - (j + lis[i]) - 1
			if best > tmp {
				best = tmp
			}
		}
		if j == p {
			stack[p] = nums[i]
			p++
		} else {
			stack[j] = nums[i]
		}
	}

	return best
}

func search(n int, fn func(int) bool) int {
	var left, right = 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
