package main

func checkRecord(n int) int {
	//no A: f(n) = f(n- 1) + f(n - 2) + f(n - 3)
	//has one A: change any one of letters from the above to A n * f(n)
	//f(1) = 1
	if n == 1 {
		return 3
	}
	if n == 2 {
		return 8
	}

	if n == 3 {
		return 19
	}

	MOD := int64(1000000007)

	f := make([]int64, n+1)
	f[1] = 2 //P & L
	f[2] = 4
	f[3] = 7

	for i := 4; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % MOD
	}

	ans := f[n] //no A

	for i := 0; i < n; i++ {
		// put a at position i
		tmp := int64(0)
		if i == 0 || i == n-1 {
			tmp = f[n-1]
		} else {
			tmp = (f[i] * f[n-i-1]) % MOD
		}
		ans = (ans + tmp) % MOD
	}

	return int(ans)
}
