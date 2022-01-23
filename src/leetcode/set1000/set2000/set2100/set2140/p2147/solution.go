package p2147

const MOD = 1000000007

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func numberOfWays(corridor string) int {
	n := len(corridor)
	// 如果seat的数目不能整除2，那么结果为0
	res := int64(1)
	prev := -1
	for i := 0; i < n; {
		// find first two seats
		var cnt int
		first := -1
		for i < n && cnt < 2 {
			if corridor[i] == 'S' {
				cnt++
				if first < 0 {
					first = i
				}
			}
			i++
		}

		if cnt == 0 {
			// last part
			if prev > 0 {
				break
			}
		}

		if cnt < 2 {
			// no two seats
			return 0
		}

		// first is first seat, i is second
		if prev >= 0 {
			res *= int64(first - prev + 1)
			res %= MOD
		}
		prev = i
	}

	return int(res)
}
