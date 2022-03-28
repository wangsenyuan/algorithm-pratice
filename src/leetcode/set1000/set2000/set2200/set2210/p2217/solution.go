package p2217

func kthPalindrome(queries []int, intLength int) []int64 {
	h := intLength / 2
	d := h + intLength&1
	cnt := make([][]int, d)
	for i := 0; i < d; i++ {
		cnt[i] = make([]int, 10)
	}

	pw := make([]int, d+1)
	pw[0] = 1
	var base int64 = 1
	for i := 0; i < h; i++ {
		base *= 10
		pw[i+1] = 10 * pw[i]
	}

	getKth := func(k int) int64 {
		var res int64
		for i := d - 1; i >= 0 && k > 0; i-- {
			var x int
			if i == d-1 {
				x++
			}
			y := x
			for x <= 9 && pw[i]*(x-y+1) < k {
				x++
			}

			if x > 9 || pw[i]*(x-y+1) < k {
				return -1
			}
			// sum[i-1] * x <= k
			res = res*10 + int64(x)
			k -= pw[i] * (x - y)
		}
		if k <= 0 {
			return -1
		}

		if intLength&1 == 0 {
			return res*base + reverse(res)
		}
		return res*base + reverse(res/10)
	}

	ans := make([]int64, len(queries))

	for i, q := range queries {
		ans[i] = getKth(q)
	}

	return ans
}

func reverse(num int64) int64 {
	var res int64
	for num > 0 {
		res = res*10 + num%10
		num /= 10
	}
	return res
}
