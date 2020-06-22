package p1488

func avoidFlood(rains []int) []int {
	n := len(rains)

	arr := make([]int, 2*n+3)

	for i := 0; i < n; i++ {
		arr[i+n] = i
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min(arr[i<<1], arr[i<<1|1])
	}

	update := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		l += n
		r += n
		res := n + 1

		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	ans := make([]int, n)
	lastRain := make(map[int]int)
	for i := 0; i < n; i++ {
		ans[i] = 1

		if rains[i] == 0 {
			continue
		}
		ans[i] = -1
		lake := rains[i]

		if k, found := lastRain[lake]; found {
			//it is full,  need to empty
			j := get(k+1, i)
			if j > i {
				return nil
			}
			ans[j] = lake

			update(j, n+1)
		}

		update(i, n+1)
		lastRain[lake] = i
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
