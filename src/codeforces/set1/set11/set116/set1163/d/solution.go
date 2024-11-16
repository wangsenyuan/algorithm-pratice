package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	c := readString(reader)
	s := readString(reader)
	t := readString(reader)
	res := solve(c, s, t)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const inf = 1 << 60

func solve(c string, s string, t string) int {
	// n := len(c)
	m := len(s)
	k := len(t)
	dp := make([][]int, m+1)
	fp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, k+1)
		fp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -inf
			fp[i][j] = -inf
		}
	}
	dp[0][0] = 0

	ps := kmp(s)
	pt := kmp(t)

	process := func(pos int, x byte) {
		for u := 0; u <= m; u++ {
			for v := 0; v <= k; v++ {
				if dp[u][v] <= -inf {
					continue
				}
				score := dp[u][v]
				nu, nv := u, v

				for nu > 0 && x != s[nu] {
					nu = ps[nu-1]
				}

				if x == s[nu] {
					nu++
					if nu == m {
						score++
						nu = ps[m-1]
					}
				}
				for nv > 0 && x != t[nv] {
					nv = pt[nv-1]
				}
				if x == t[nv] {
					nv++
					if nv == k {
						score--
						nv = pt[k-1]
					}
				}

				fp[nu][nv] = max(fp[nu][nv], score)
			}
		}
	}

	for i, x := range []byte(c) {
		if x != '*' {
			process(i, x)
		} else {
			for y := byte('a'); y <= 'z'; y++ {
				process(i, y)
			}
		}
		for u := 0; u < m; u++ {
			for v := 0; v < k; v++ {
				dp[u][v] = fp[u][v]
				fp[u][v] = -inf
			}
		}
	}
	res := -inf
	for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			res = max(res, dp[i][j])
		}
	}
	return res
}

func kmp(p string) []int {
	n := len(p)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		j := res[i-1]
		for j > 0 && p[i] != p[j] {
			j = res[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		res[i] = j
	}
	return res
}
