package lcp80

import (
	"bytes"
	"sort"
)

func evolutionaryRecord(parents []int) string {
	// 考虑两个child，是按照深度排序，还是按照size呢？
	// deeper，0的位置越靠前
	// 不是简单的这个比较
	// 考虑一个中间的节点u，它一共有sz[u]的子节点
	// 每个子节点贡献一个0，但是共享 sz[u] - leaf[u]个1
	// 结果必然是 {a}{b}{c}{d}
	// 两个字符串比较的方式是 x0 > y0 or x0 == y0 && x1 < y1 || x0 == y0 && x1 == y1 && x2 > y2
	// 这里缺少一个信息，如何确定，u子树的字符串比v子树的小？
	// 高度吗？这个应该是相关的，但似乎不是唯一的
	// size吗？目前还不明确
	// 是同一层上面，节点的个数吗？
	// 000101000111
	// 似乎是的
	// 怎么实现呢？
	//
	n := len(parents)

	if n == 1 {
		return ""
	}

	g := NewGraph(n, n)
	var root int
	for i, p := range parents {
		if p >= 0 {
			g.AddEdge(p, i)
		} else {
			root = i
		}
	}

	var dfs func(u int) string

	dfs = func(u int) string {
		var arr []string
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			tmp := "0" + dfs(v) + "1"
			arr = append(arr, tmp)
		}
		sort.Strings(arr)
		var buf bytes.Buffer
		for i := 0; i < len(arr); i++ {
			buf.WriteString(arr[i])
		}
		return buf.String()
	}

	res := dfs(root)
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == '1' {
			continue
		}
		return res[:i+1]
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
