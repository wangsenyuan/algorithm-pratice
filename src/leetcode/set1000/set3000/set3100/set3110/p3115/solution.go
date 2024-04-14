package p3115

func findKthSmallest(coins []int, k int) int64 {
	n := len(coins)

	count := func(expect int) int {
		var res int
		for state := 1; state < 1<<n; state++ {
			var cnt int
			l := 1
			for i := 0; i < n && l <= expect; i++ {
				if (state>>i)&1 == 1 {
					l = lcm(l, coins[i])
					cnt++
				}
			}
			if l > expect {
				continue
			}
			tmp := expect / l
			if cnt&1 == 1 {
				res += tmp
			} else {
				res -= tmp
			}
		}
		return res
	}

	l, r := 0, 1<<62

	for l < r {
		mid := (l + r) / 2
		if count(mid) >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return int64(r)
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
