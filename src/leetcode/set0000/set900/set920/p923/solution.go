package p923

const MAX_TARGE = 300
const MAX_NUM = 100
const MOD = 1e9 + 7

func threeSumMulti(A []int, target int) int {
	cnt := make([]int64, MAX_NUM+1)
	n := len(A)
	for i := 0; i < n; i++ {
		cnt[A[i]]++
	}

	sum1 := func() int {
		if target%3 != 0 {
			return 0
		}
		num := target / 3
		nc := cnt[num]
		ans := (nc * (nc - 1) * (nc - 2) / 6) % MOD
		return int(ans)
	}

	sum2 := func() int {
		var ans int64
		for a := 0; a <= MAX_NUM; a++ {
			if cnt[a] < 2 {
				continue
			}
			b := target - 2*a
			if b < 0 || b > MAX_NUM || cnt[b] == 0 || a == b {
				continue
			}
			// choose 2 a, and one b
			res := (cnt[a] * (cnt[a] - 1) / 2) % MOD
			res = (res * int64(cnt[b])) % MOD
			ans += res
			if ans >= MOD {
				ans -= MOD
			}
		}
		return int(ans)
	}

	sum3 := func() int {
		var ans int64

		for a := 0; a <= MAX_NUM; a++ {
			if cnt[a] == 0 {
				continue
			}
			for b := a + 1; b <= MAX_NUM; b++ {
				if cnt[b] == 0 {
					continue
				}
				c := target - b - a
				if c < 0 || c > MAX_NUM || c <= b || cnt[c] == 0 {
					continue
				}
				res := (cnt[a] * cnt[b] * cnt[c]) % MOD
				ans += res
				if ans >= MOD {
					ans -= MOD
				}
			}
		}
		return int(ans)
	}

	res := sum1()
	res += sum2()
	if res >= MOD {
		res -= MOD
	}
	res += sum3()
	if res >= MOD {
		res -= MOD
	}
	return res
}
