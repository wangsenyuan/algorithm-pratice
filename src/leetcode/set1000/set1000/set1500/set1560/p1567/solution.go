package p1567

func getMaxLen(nums []int) int {
	merge := func(left, right, mid int) int {
		if nums[mid] == 0 {
			return 0
		}
		pos := make([]int, 2)
		var prod = 1
		for i := mid; i < right && nums[i] != 0; i++ {
			if nums[i] < 0 {
				prod *= -1
			}
			if prod > 0 {
				pos[0] = i - mid + 1
			} else {
				pos[1] = i - mid + 1
			}
		}
		var best int
		prod = 1
		for i := mid - 1; i >= left && nums[i] != 0; i-- {
			if nums[i] < 0 {
				prod *= -1
			}
			if prod > 0 {
				best = max(best, mid-i+pos[0])
			} else if pos[1] > 0 {
				best = max(best, mid-i+pos[1])
			}
		}

		return best
	}

	var loop func(i, j int) int

	loop = func(i, j int) int {
		if i+1 == j {
			if nums[i] > 0 {
				return 1
			}
			return 0
		}
		mid := (i + j) / 2
		a := loop(i, mid)
		b := loop(mid, j)
		c := merge(i, j, mid)

		return max3(a, b, c)
	}

	return loop(0, len(nums))
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
