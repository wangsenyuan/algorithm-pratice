package p2571

const MOD = 1e9 + 7

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

func squareFreeSubsets(nums []int) int {
	var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	// n <= 100
	//
	// nums[i] <= 30
	var n int
	var cnt1 int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt1++
		} else {
			nums[n] = nums[i]
			n++
		}
	}
	if n == 0 {
		return modSub(pow(2, cnt1), 1)
	}
	m := len(primes)
	// prod 最多只能整除primes[i] 一次
	M := 1 << m
	dp := make([]int, M)
	dp[0] = 1
	fp := make([]int, M)
	for i := 0; i < n; i++ {
		tmp := nums[i]
		ok := true
		var cur int
		for j := 0; j < m && ok; j++ {
			if tmp%primes[j] == 0 {
				var cnt int
				for tmp%primes[j] == 0 {
					cnt++
					tmp /= primes[j]
				}
				ok = cnt == 1
				cur |= (1 << j)
			}
		}
		if !ok {
			continue
		}
		for j := 0; j < M; j++ {
			fp[j] = 0
		}
		for state := 0; state < M; state++ {
			if state&cur == 0 && cur > 0 {
				fp[state|cur] = modAdd(fp[state|cur], dp[state])
			}
		}
		for state := 0; state < M; state++ {
			dp[state] = modAdd(dp[state], fp[state])
		}
	}

	var res int

	for state := 1; state < M; state++ {
		res = modAdd(res, dp[state])
	}
	x1 := pow(2, cnt1)
	return modAdd(modMul(res, x1), modSub(x1, 1))
}
