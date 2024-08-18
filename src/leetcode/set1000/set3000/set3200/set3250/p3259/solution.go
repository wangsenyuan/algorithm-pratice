package p3259

import (
	"slices"
)

func largestPalindrome(n int, mod int) (ans string) {
	if n == 1 {
		for i := 9; ; i-- {
			if i%mod == 0 {
				return string('0' + byte(i))
			}
		}
	}
	pow10 := make([]int, n)
	pow10[0] = 1
	for i := 1; i < n; i++ {
		pow10[i] = pow10[i-1] * 10 % mod
	}

	m := (n + 1) / 2
	dp := make([][]int8, m)
	for i := range dp {
		dp[i] = make([]int8, mod)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	type pair struct{ to, d int }
	from := make([][]pair, m)
	for i := range from {
		from[i] = make([]pair, mod)
		for j := range from[i] {
			from[i][j].to = -1
		}
	}

	var f func(int, int) int8
	f = func(i, j int) (res int8) {
		if i == m {
			if j == 0 {
				return 1
			}
			return 0
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for k := 9; k >= 0; k-- {
			var j2 int
			if n%2 == 1 && i == m-1 {
				j2 = (j + k*pow10[i]) % mod
			} else {
				j2 = (j + k*(pow10[i]+pow10[n-1-i])) % mod
			}
			r := f(i+1, j2)
			if r == 1 {
				from[i][j] = pair{j2, k}
				return 1
			}
		}
		return 0
	}
	for d := 9; d > 0; d-- {
		modD := d * (1 + pow10[n-1]) % mod
		if f(1, modD) == 1 {
			path := []byte{'0' + byte(d)}
			i, j := 1, modD
			for i < m {
				path = append(path, '0'+byte(from[i][j].d))
				j = from[i][j].to
				i++
			}

			s := string(path)
			_tmp := []byte(s)
			slices.Reverse(_tmp)
			rev := string(_tmp)
			if n%2 == 1 {
				rev = rev[1:]
			}
			s += rev
			ans = s
			break
		}
	}

	return
}
