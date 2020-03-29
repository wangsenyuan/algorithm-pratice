package p1396

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	m := len(evil)
	next := kmp(evil)

	dp := make([][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, 4)
		for j := 0; j < 4; j++ {
			dp[i][j] = make([]int64, m)
		}
	}
	dp[0][3][0] = 1

	for i := 0; i < n; i++ {
		for d := 0; d < 4; d++ {
			d1 := d & 1
			d2 := (d & 2)
			var start int
			if d1 == 1 {
				start = int(s1[i] - 'a')
			}
			var end int = 25
			if d2 == 2 {
				end = int(s2[i] - 'a')
			}

			for x := start; x <= end; x++ {
				var dd1 int
				if d1 == 1 && x == start {
					dd1 = 1
				}
				var dd2 int
				if d2 == 2 && x == end {
					dd2 = 2
				}
				for k := 0; k < m; k++ {
					kk := k
					for kk > 0 && x != int(evil[kk]-'a') {
						kk = next[kk-1]
					}

					if x == int(evil[kk]-'a') {
						kk++
					}
					if kk < m {
						dp[i+1][dd1|dd2][kk] += dp[i][d][k]
						if dp[i+1][dd1|dd2][kk] >= MOD {
							dp[i+1][dd1|dd2][kk] -= MOD
						}
					}
				}
			}
		}

	}

	var res int64

	for d := 0; d < 4; d++ {
		for k := 0; k < m; k++ {
			res += dp[n][d][k]
			if res >= MOD {
				res -= MOD
			}
		}
	}

	return int(res)
}

func findGoodStrings3(n int, s1 string, s2 string, evil string) int {
	m := len(evil)
	next := kmp(evil)

	dp := make([][][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]int64, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([][]int64, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = make([]int64, m)
			}
		}
	}
	dp[0][1][1][0] = 1

	for i := 0; i < n; i++ {
		for d1 := 0; d1 < 2; d1++ {
			for d2 := 0; d2 < 2; d2++ {
				var start int
				if d1 == 1 {
					start = int(s1[i] - 'a')
				}
				var end int = 25
				if d2 == 1 {
					end = int(s2[i] - 'a')
				}

				for x := start; x <= end; x++ {
					var dd1 int
					if d1 == 1 && x == start {
						dd1 = 1
					}
					var dd2 int
					if d2 == 1 && x == end {
						dd2 = 1
					}
					for k := 0; k < m; k++ {
						kk := k
						for kk > 0 && x != int(evil[kk]-'a') {
							kk = next[kk-1]
						}

						if x == int(evil[kk]-'a') {
							kk++
						}
						if kk < m {
							dp[i+1][dd1][dd2][kk] += dp[i][d1][d2][k]
							dp[i+1][dd1][dd2][kk] %= MOD
						}
					}
				}

			}
		}
	}

	var res int64

	for d1 := 0; d1 < 2; d1++ {
		for d2 := 0; d2 < 2; d2++ {
			for k := 0; k < m; k++ {
				res += dp[n][d1][d2][k]
				res %= MOD
			}
		}
	}

	return int(res)
}

func findGoodStrings2(n int, s1 string, s2 string, evil string) int {
	m := len(evil)
	dp := make([][][][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][][]int64, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([][]int64, 2)
			for k := 0; k < 2; k++ {
				dp[i][j][k] = make([]int64, m)
				for a := 0; a < m; a++ {
					dp[i][j][k][a] = -1
				}
			}
		}
	}

	next := kmp(evil)

	var loop func(i int, preS1 int, preS2 int, preEvil int) int64

	loop = func(i int, prevS1 int, prevS2 int, preEvil int) int64 {
		if preEvil == m {
			return 0
		}
		if i == n {
			return 1
		}
		if dp[i][prevS1][prevS2][preEvil] >= 0 {
			return dp[i][prevS1][prevS2][preEvil]
		}
		dp[i][prevS1][prevS2][preEvil] = 0

		var start int
		if prevS1 == 1 {
			start = int(s1[i] - 'a')
		}

		var last int = int('z' - 'a')
		if prevS2 == 1 {
			last = int(s2[i] - 'a')
		}

		for x := start; x <= last; x++ {
			var _prevS1 int
			if prevS1 == 1 && x == start {
				_prevS1 = 1
			}
			var _prevS2 int
			if prevS2 == 1 && x == last {
				_prevS2 = 1
			}

			_prevEvil := preEvil

			for _prevEvil > 0 && x != int(evil[_prevEvil]-'a') {
				_prevEvil = next[_prevEvil-1]
			}
			if x == int(evil[_prevEvil]-'a') {
				_prevEvil++
			}

			dp[i][prevS1][prevS2][preEvil] += loop(i+1, _prevS1, _prevS2, _prevEvil)
			dp[i][prevS1][prevS2][preEvil] %= MOD
		}

		return dp[i][prevS1][prevS2][preEvil]
	}

	return int(loop(0, 1, 1, 0))
}

func kmp(s string) []int {
	res := make([]int, len(s))
	var l int

	for i := 1; i < len(s); i++ {
		for l > 0 && s[l] != s[i] {
			l = res[l-1]
		}
		if s[l] == s[i] {
			l++
		}
		res[i] = l
	}
	return res
}

const MOD = 1000000007

func findGoodStrings1(n int, s1 string, s2 string, evil string) int {
	// dp[i][j] is string s with lengh i,
	// n := len(s1)
	a, _ := find(n, s1, evil)
	_, d := find(n, s2, evil)

	return (d - a + MOD) % MOD
}

func find(n int, s string, evil string) (int, int) {
	m := len(evil)

	next := kmp(evil)

	// dp[i][j] is a string with length i, and match prev j evil string

	dp := make([][][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([][]int, 26)
			for k := 0; k < 26; k++ {
				dp[i][j][k] = make([]int, m)
			}
		}
	}

	for x := 0; x < 26; x++ {
		if x > int(s[0]-'a') {
			break
		}
		var d int
		if x == int(s[0]-'a') {
			d = 1
		}
		var j int
		if x == int(evil[0]-'a') {
			j = 1
		}
		if j < m {
			dp[0][d][x][j] = 1
		}
	}

	for i := 0; i < n-1; i++ {
		// try to put byte x at place i - 1
		for d := 0; d <= 1; d++ {
			for x := 0; x < 26; x++ {
				for y := 0; y < 26; y++ {
					if d == 1 && y > int(s[i+1]-'a') {
						continue
					}
					var dd int

					if d == 1 && y == int(s[i+1]-'a') {
						dd = 1
					}

					for k := 0; k < m; k++ {
						kk := k
						for kk > 0 && y != int(evil[kk]-'a') {
							kk = next[kk-1]
						}
						if y == int(evil[kk]-'a') {
							kk++
						}
						if kk < m {
							dp[i+1][dd][y][kk] += dp[i][d][x][k]
							dp[i+1][dd][y][kk] %= MOD
						}
					}
				}
			}
		}
	}
	var first int
	var second int

	for l := 0; l < m; l++ {
		for x := 0; x < 26; x++ {
			first += dp[n-1][0][x][l]
			first %= MOD
			second += dp[n-1][1][x][l]
			second %= MOD
		}
	}
	second += first
	second %= MOD
	return first, second
}
