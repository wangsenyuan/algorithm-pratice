package p2943

func minimumCoins(prices []int) int {
	n := len(prices)
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 << 30
	}

	set := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 1 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		res := 1 << 30
		l += n
		r += n
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

	for i := n - 1; i >= 0; i-- {
		j := i + 1
		if j+j >= n {
			// get all after prices[i]
			set(i, prices[i])
		} else {
			// j + j < n
			tmp := get(i+1, j+j+1)
			set(i, prices[i]+tmp)
		}
	}

	return get(0, 1)
}
