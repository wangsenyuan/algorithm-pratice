package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	t := readString(reader)
	res := solve(s, t)

	a := fmt.Sprintf("%v", res)

	fmt.Println(a[1 : len(a)-1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func solve(s string, t string) []int {
	mp := ['g']int{}

	n := len(s)
	m := len(t)
	ans := make([]int, n-m+1)

	next := make([]int, m)

	process := func(sz int) {
		clear(next)
		for i := 1; i < m; i++ {
			j := next[i-1]
			for j > 0 && mp[t[i]] != mp[t[next[j-1]]] {
				j = next[j-1]
			}
			if mp[t[i]] == mp[t[j]] {
				j++
			}

			next[i] = j
		}

		for i, j := 0, 0; i < n; i++ {
			for j > 0 && mp[s[i]] != mp[t[j]] {
				j = next[j-1]
			}
			if mp[s[i]] == mp[t[j]] {
				j++
			}
			if j == m {
				ans[i-m+1] = max(ans[i-m+1], sz)
				j = next[j-1]
			}
		}
	}

	var dfs func(c int, sz int)

	dfs = func(c int, sz int) {
		if c == 'g' {
			process(sz)
			return
		}
		mp[c] = sz
		dfs(c+1, sz+1)
		for mp[c] = range sz {
			dfs(c+1, sz)
		}
	}

	dfs('a', 0)

	for i := range len(ans) {
		ans[i] = 6 - ans[i]
	}

	return ans
}
