package p3021

func maximumLength(nums []int) int {
	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}

	check := func(x int) int {
		// x is the first
		if x == 1 {
			if cnt[1]%2 == 0 {
				return cnt[1] - 1
			}
			return cnt[1]
		}
		// x > 1
		var res int
		for cnt[x] >= 2 {
			res += 2
			x *= x
		}
		if cnt[x] == 1 {
			res++
		} else {
			res--
		}
		return res
	}

	best := 1
	for _, num := range nums {
		best = max(best, check(num))
	}
	return best
}
