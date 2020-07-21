package p1520

import (
	"sort"
)

func maxNumOfSubstrings(s string) []string {
	last := make([]int, 26)
	first := make([]int, 26)
	for i := 0; i < len(last); i++ {
		last[i] = -1
		first[i] = -1
	}

	n := len(s)

	for i := 0; i < n; i++ {
		j := ind(s[i])
		if first[j] < 0 {
			first[j] = i
		}
		last[j] = i
	}
	items := make(Items, 0, 26)

	for i := 0; i < 26; i++ {
		if first[i] < 0 {
			continue
		}
		l, r := first[i], last[i]
		keep := true
		for k := l; k <= r; k++ {
			j := ind(s[k])
			if first[j] < l || last[j] > r {
				if first[j] < l {
					keep = false
					break
				}
				r = max(r, last[j])
			}
		}
		if keep {
			items = append(items, Item{l, r})
		}
	}

	sort.Sort(items)
	var res []string
	var prev = -1
	for i := 0; i < len(items); i++ {
		if items[i].first < prev {
			continue
		}
		res = append(res, s[items[i].first:items[i].second+1])
		prev = items[i].second
	}

	return res
}

type Item struct {
	first, second int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func width(item Item) int {
	return item.second - item.first
}

func (items Items) Less(i, j int) bool {
	return items[i].second < items[j].second || items[i].second == items[j].second && width(items[i]) < width(items[j])
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func maxNumOfSubstrings1(s string) []string {
	last := make([]int, 26)
	first := make([]int, 26)
	for i := 0; i < len(last); i++ {
		last[i] = -1
		first[i] = -1
	}

	conn := make([][]bool, 26)
	rConn := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		conn[i] = make([]bool, 26)
		rConn[i] = make([]bool, 26)
	}

	n := len(s)

	for i := 0; i < n; i++ {
		j := ind(s[i])
		if first[j] < 0 {
			first[j] = i
		}

		for k := 0; k < 26; k++ {
			if j != k && first[j] < last[k] {
				conn[j][k] = true
			}
		}
		last[j] = i
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			rConn[i][j] = conn[j][i]
		}
	}

	res := make([]string, 0, 26)

	writeResult := func(begin, end int) {
		if end < 0 {
			return
		}

		res = append(res, s[begin:end+1])
	}

	// 使用两次DFS找出图中的SCC
	// 第一次POST-ORDER DFS作用在原来的图上，将节点放入stack中
	// 第二次DFS从stack中pop出节点，在翻转的图上，将SCC标示出来，并且计算该SCC的范围，如果有其他SCC的边进入该SCC，则标示它为非叶子节点；
	// 最后把所有叶子节点输出出来

	que := make([]int, 0, 26)
	vis := make([]bool, 26)

	var dfs1 func(u int)

	dfs1 = func(u int) {
		vis[u] = true
		for v := 0; v < 26; v++ {
			if conn[u][v] && !vis[v] {
				dfs1(v)
			}
		}
		que = append(que, u)
	}

	for i := 0; i < 26; i++ {
		if first[i] >= 0 && !vis[i] {
			dfs1(i)
		}
	}

	leaf := make([]bool, 26)

	comp := make([]int, 26)
	for i := 0; i < 26; i++ {
		comp[i] = -1
	}
	var dfs2 func(u int, c int, rg []int)

	dfs2 = func(u int, c int, rg []int) {
		comp[u] = c
		rg[0] = min(rg[0], first[u])
		rg[1] = max(rg[1], last[u])
		for v := 0; v < 26; v++ {
			if rConn[u][v] {
				if comp[v] < 0 {
					dfs2(v, c, rg)
				} else if comp[v] != c {
					leaf[comp[v]] = false
				}
			}
		}
	}

	for i := len(que) - 1; i >= 0; i-- {
		u := que[i]
		if comp[u] < 0 {
			// u is also the SCC component number
			leaf[u] = true
			rg := []int{n, -1}
			dfs2(u, u, rg)
			first[u] = rg[0]
			last[u] = rg[1]
		}
	}

	for i := len(que) - 1; i >= 0; i-- {
		u := que[i]
		if leaf[u] {
			writeResult(first[u], last[u])
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func ind(b byte) int {
	return int(b - 'a')
}
