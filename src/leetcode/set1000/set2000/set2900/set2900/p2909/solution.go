package p2909

func minGroupsForValidAssignment(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	var arr []int
	x := len(nums)
	for _, v := range freq {
		arr = append(arr, v)
		x = min(x, v)
	}

	check := func(d int) int {
		// 每组是d或者d-1个
		var cnt int
		for _, v := range arr {
			a, b := v/d, v%d
			if b == 0 {
				cnt += a
				continue
			}
			if b == d-1 {
				cnt += a + 1
				continue
			}
			if d-1-b > a {
				return -1
			}
			cnt += a + 1
		}
		return cnt
	}

	best := len(nums)
	for i := 1; i <= x+1; i++ {
		// 每组最多i个
		tmp := check(i)
		if tmp > 0 {
			best = min(best, tmp)
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
