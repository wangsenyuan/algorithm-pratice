package p2217

func kthPalindrome(queries []int, intLength int) []int64 {
	h := intLength / 2
	d := h + intLength&1
	cnt := make([][]int, d)
	for i := 0; i < d; i++ {
		cnt[i] = make([]int, 10)
	}
	sum := make([]int, d)
	for i := 0; i < d; i++ {
		for x := 0; x < 10; x++ {
			if i > 0 {
				cnt[i][x] = sum[i-1]
			} else {
				cnt[i][x] = 1
			}
		}
		for x := 0; x < 10; x++ {
			sum[i] += cnt[i][x]
			if x > 0 {
				cnt[i][x] += cnt[i][x-1]
			}
		}
	}

	var base int64 = 1

	for i := 0; i < h; i++ {
		base *= 10
	}

	getSum := func(i int) int {
		if i < 0 {
			return 1
		}
		return sum[i]
	}

	getKth := func(k int) int64 {
		var res int64
		for i := d - 1; i >= 0 && k > 0; i-- {
			var x int
			if i == d-1 {
				x++
			}
			y := x
			for x <= 9 && getSum(i-1)*(x-y+1) < k {
				x++
			}

			if x > 9 || getSum(i-1)*(x-y+1) < k {
				return -1
			}
			// sum[i-1] * x <= k
			res = res*10 + int64(x)
			k -= getSum(i-1) * (x - y)
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
