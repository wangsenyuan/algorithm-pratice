package p3327

import "bytes"

func findAnswer(parent []int, s string) []bool {
	n := len(parent)
	g := make([][]int, n)

	for i := 1; i < n; i++ {
		g[parent[i]] = append(g[parent[i]], i)
	}

	var dfs func(u int)

	buf := make([]byte, n)
	var it int
	sz := make([]int, n)
	pos := make([]int, n)
	dfs = func(u int) {
		sz[u]++
		for _, v := range g[u] {
			dfs(v)
			sz[u] += sz[v]
		}
		buf[it] = s[u]
		pos[u] = it
		it++
	}

	dfs(0)

	f := manachers(modify(string(buf)))
	ans := make([]bool, n)
	for u := 0; u < n; u++ {
		i := pos[u]*2 + 1
		j := 2*(pos[u]-sz[u]+1) + 1
		ans[u] = f[(i+j)/2] >= sz[u]
	}

	return ans
}

func modify(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		buf.WriteByte('#')
		buf.WriteByte(s[i])
	}
	buf.WriteByte('#')
	return buf.String()
}

func manachers(s string) []int {
	s = "&" + s + "^"
	n := len(s)
	p := make([]int, n)

	l, r := 1, 1

	for i := 1; i < n-1; i++ {
		p[i] = max(0, min(r-i, p[l+(r-i)]))
		for s[i-p[i]] == s[i+p[i]] {
			p[i]++
		}
		if i+p[i] > r {
			l = i - p[i]
			r = i + p[i]
		}
	}

	return p[1 : n-1]
}
