package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)
	s = fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const inf = 1 << 30

func solve(s string, t string) []int {
	lp := kmp(t)
	n := len(s)
	m := len(t)
	k := n - m + 1
	dp := make([][]int, k+1)
	fp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, m+1)
		fp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = inf
			fp[i][j] = inf
		}
	}

	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = inf
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for u := 0; u <= k; u++ {
			for v := 0; v < m; v++ {
				if dp[u][v] == inf {
					continue
				}
				// 这一段是为了能够匹配，并增加u
				if s[i] == t[v] {
					nu, nv := u, v+1
					if nv == m {
						nu++
						nv = lp[m-1]
					}
					if nu <= k {
						fp[nu][nv] = min(fp[nu][nv], dp[u][v])
					}
					// 或者通过修改当前字符和s[i]不同
					nu, nv = u, v
					// 这个状态感觉不大对，
					for nv > 0 && t[nv] == s[i] {
						nv = lp[nv-1]
					}
					if t[nv] != s[i] {
						nv++
					}
					// 通过修改t[v] = 1 ^ t[v]
					fp[nu][nv] = min(fp[nu][nv], dp[u][v]+1)
				} else {
					// 修改成和s[i]相同的字符
					nu, nv := u, v+1
					if nv == m {
						nu++
						nv = lp[m-1]
					}
					if nu <= k {
						fp[nu][nv] = min(fp[nu][nv], dp[u][v]+1)
					}
					// 或者不修改
					nu, nv = u, v
					for nv > 0 && t[nv] != s[i] {
						nv = lp[nv-1]
					}
					if t[nv] == s[i] {
						nv++
					}
					// 通过修改t[v] = 1 ^ t[v]
					fp[nu][nv] = min(fp[nu][nv], dp[u][v])
				}
			}
		}
		for u := 0; u <= k; u++ {
			for v := 0; v < m; v++ {
				dp[u][v] = fp[u][v]
				fp[u][v] = inf
				// 把i后面的都要删除掉
				res[u] = min(res[u], dp[u][v]+n-1-i)
			}
		}
	}

	for i := 0; i <= k; i++ {
		if res[i] == inf {
			res[i] = -1
		}
	}

	return res
}

func kmp(p string) []int {
	n := len(p)
	next := make([]int, n)
	for i := 1; i < n; i++ {
		j := next[i-1]
		for j > 0 && p[i] != p[j] {
			j = next[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		next[i] = j
	}
	return next
}
