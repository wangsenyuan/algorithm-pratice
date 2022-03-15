package c

import "sort"

const MOD = 1000000007

//2999992273
func maxInvestment(product []int, limit int) int {
	sort.Ints(product)
	var best int
	var sum int64

	for _, price := range product {
		sum += int64(price)
	}

	n := len(product)

	profit := func(s int, low int) int64 {
		var res int64
		for i := s; i < n; i++ {
			tmp := (int64(low+product[i]) * int64(product[i]-low+1) / 2) % MOD
			res += tmp
			res %= MOD
		}
		return res
	}

	if sum <= int64(limit) {
		return int(profit(0, 1))
	}

	for i := 0; i < n; i++ {
		// if keep some value of current until product[i]
		// sum - x * (n - i) <= limit
		if sum >= int64(limit) {
			x := (sum - int64(limit)) / int64(n-i)
			if x+1 <= int64(product[i]) {
				var extra int64
				if sum-x*int64(n-i) > int64(limit) {
					// need to remove more x
					x++
					// extra = x * (int64(limit) - (sum - x*int64(n-i))) % MOD
					extra = mul(x, sub(int64(limit), sub(sum, mul(x, int64(n-i)))))
				}
				if x+1 <= int64(product[i]) {
					return int(add(profit(i, int(x+1)), extra))
				}
			}
		}
		sum -= int64(product[i])
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func add(a, b int64) int64 {
	a += b
	return a % MOD
}

func sub(a, b int64) int64 {
	return add(a, MOD-b)
}

func mul(a, b int64) int64 {
	a *= b
	a %= MOD
	return a
}
