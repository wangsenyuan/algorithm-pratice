package p2584

func findValidSplit(nums []int) int {
	n := len(nums)
	// A[i] <= 1e6
	x := max_of(nums)
	spf := make([]int, x+1)
	for i := 2; i <= x; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if x/i < i {
				continue
			}
			for j := i * i; j <= x; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	last_occ := make(map[int]int)
	next := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		next[i] = i
		a := nums[i]
		for a > 1 {
			b := spf[a]
			if _, ok := last_occ[b]; ok {
				next[i] = max(next[i], last_occ[b])
			}
			a /= b
		}

		a = nums[i]

		for a > 1 {
			b := spf[a]
			if _, ok := last_occ[b]; !ok {
				last_occ[b] = i
			}
			a /= b
		}
	}

	p := next[0]

	for i := 1; i < n; i++ {
		if p >= i {
			p = max(p, next[i])
		} else {
			return i - 1
		}
	}

	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max_of(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
