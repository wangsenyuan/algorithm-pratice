package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	return
}

func readTripNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	readInt(bs, x+1, &c)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		p, q, n := readTripNums(scanner)
		b1, b2 := readTwoNums(scanner)
		s1, s2 := readTwoNums(scanner)
		fmt.Println(solve(p, q, n, b1, b2, s1, s2))
	}
}

const MOD = 1000000007

func solve1(P, Q, N, B1, B2, S1, S2 int) int {
	f := func(S int, dp [][]int) func(int, int) int {
		var fn func(j int, k int) int

		fn = func(j int, k int) int {
			if j < 0 || k < 0 {
				return 0
			}
			if j >= len(dp) || k >= len(dp[j]) {
				return 0
			}

			if dp[j][k] == -1 {
				var tmp int
				for i := 1; i <= S; i++ {
					tmp = (tmp + fn(j-1, k-i)) % MOD
				}
				dp[j][k] = tmp
			}

			return dp[j][k]
		}

		return fn
	}

	dp1 := make([][]int, B1+1)
	for i := 0; i <= B1; i++ {
		dp1[i] = make([]int, S1+1)
		for j := 0; j <= S1; j++ {
			dp1[i][j] = -1
		}
	}
	dp1[0][0] = 1
	fa := f(S1, dp1)
	dp2 := make([][]int, B2+1)
	for i := 0; i <= B2; i++ {
		dp2[i] = make([]int, S2+1)
		for j := 0; j <= S2; j++ {
			dp2[i][j] = -1
		}
	}
	dp2[0][0] = 1
	fb := f(S2, dp2)

	var ans int

	for l := 1; l <= N; l++ {
		m := l / 2
		n := l - m
		for x := 0; x <= N; x++ {
			tmp := (fa(m, x) * fb(n, N-x)) % MOD
			tmp = (tmp + fa(n, x)*fb(m, N-x)) % MOD
			ans = (ans + tmp) % MOD
		}
	}
	return ans
}

func solve(P, Q, N, B1, B2, S1, S2 int) int {
	dp := make([][][][][][]int, P+1)
	for i := 0; i <= P; i++ {
		dp[i] = make([][][][][]int, Q+1)
		for j := 0; j <= Q; j++ {
			dp[i][j] = make([][][][]int, B1+1)
			for k := 0; k <= B1; k++ {
				dp[i][j][k] = make([][][]int, max(S1, S2)+1)
				for m := 0; m <= max(S1, S2); m++ {
					dp[i][j][k][m] = make([][]int, 2)
					dp[i][j][k][m][0] = make([]int, 2)
					dp[i][j][k][m][0][0] = -1
					dp[i][j][k][m][0][1] = -1
					dp[i][j][k][m][1] = make([]int, 2)
					dp[i][j][k][m][1][0] = -1
					dp[i][j][k][m][1][1] = -1
				}
			}
		}
	}

	var fn func(a, b, m, b1d, cont, box, val int) int
	fn = func(a, b, m, b1d, cont, box, val int) int {
		if a > P {
			return 0
		}
		if b > Q {
			return 0
		}
		if b1d > B1 {
			return 0
		}
		var b2d, s1d, s2d int
		if box == 0 {
			b2d = b1d - 1 + val
		} else {
			b2d = b1d + val
		}

		if b2d > B2 {
			return 0
		}

		if m == N {
			return 1
		}

		if dp[a][b][b1d][cont][box][val] != -1 {
			return dp[a][b][b1d][cont][box][val]
		}

		if box == 0 {
			s1d = cont
			s2d = 0
		} else {
			s1d = 0
			s2d = cont
		}

		if s1d == S1 {
			dp[a][b][b1d][cont][box][val] = fn(a, b+1, m+1, b1d, 1, 1, val)
			return dp[a][b][b1d][cont][box][val]
		}

		if s2d == S2 {
			dp[a][b][b1d][cont][box][val] = fn(a+1, b, m+1, b1d+1, 1, 0, val)
			return dp[a][b][b1d][cont][box][val]
		}
		if s1d > 0 {
			dp[a][b][b1d][cont][box][val] = (fn(a+1, b, m+1, b1d, s1d+1, 0, val) + fn(a, b+1, m+1, b1d, 1, 1, val)) % MOD
			return dp[a][b][b1d][cont][box][val]
		}
		dp[a][b][b1d][cont][box][val] = (fn(a+1, b, m+1, b1d+1, 1, 0, val) + fn(a, b+1, m+1, b1d, s2d+1, 1, val)) % MOD
		return dp[a][b][b1d][cont][box][val]
	}
	return (fn(1, 0, 1, 1, 1, 0, 0) + fn(0, 1, 1, 0, 1, 1, 1)) % MOD
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
