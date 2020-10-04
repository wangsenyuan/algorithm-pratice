package p1573

const MOD = 1000000007

func numWays(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
	}

	if sum%3 != 0 {
		return 0
	}
	n := int64(len(s))
	if sum == 0 {
		return int(((n - 1) * (n - 2) / 2) % MOD)
	}

	sum /= 3
	var res int64 = 1
	var cnt int
	var x int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt++
		} else {
			x++
			if x == sum+1 {
				res *= int64(cnt + 1)
				res %= MOD
				x = 1
			}
			cnt = 0
		}
	}
	return int(res)
}
