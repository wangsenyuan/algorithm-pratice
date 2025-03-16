package p3485

func longestCommonPrefix(words []string, k int) []int {
	n := len(words)
	res := make([]int, n)
	if n-1 < k {
		return res
	}

	type pair struct {
		first  int
		second int
	}

	// 似乎有点问题
	// 删除一个后，大部分情况下是不会变的
	// 那么使用trie，且只保留那些 >= k的节点
	// 且每个节点上保留两个值，最大值和第二大的值
	type node struct {
		children []*node
		cnt      int
		vals     []pair // 最多保留两个即可
	}

	tr := new(node)
	tr.children = make([]*node, 26)
	for _, w := range words {
		tmp := tr
		for j := 0; j < len(w); j++ {
			tmp.cnt++
			x := int(w[j] - 'a')
			if tmp.children[x] == nil {
				tmp.children[x] = new(node)
				tmp.children[x].children = make([]*node, 26)
			}
			tmp = tmp.children[x]
		}
		tmp.cnt++
	}

	var dfs func(node *node) *node
	dfs = func(node *node) *node {
		if node == nil || node.cnt < k {
			return nil
		}

		for i, child := range node.children {
			node.children[i] = dfs(child)
		}
		return node
	}

	tr = dfs(tr)

	if tr == nil {
		return res
	}

	var dfs2 func(node *node) pair

	dfs2 = func(node *node) pair {
		if node == nil {
			return pair{-1, -1}
		}
		node.vals = []pair{{-1, -1}, {-1, -1}}
		for i, child := range node.children {
			tmp := dfs2(child)
			// tmp[0] 是最大值
			if tmp.second+1 >= node.vals[0].second {
				node.vals[1] = node.vals[0]
				node.vals[0] = pair{i, tmp.second + 1}
			} else if tmp.second+1 >= node.vals[1].second {
				node.vals[1] = pair{i, tmp.second + 1}
			}
		}
		return node.vals[0]
	}

	dfs2(tr)

	for j, w := range words {
		tmp := tr
		for i := range len(w) {
			x := int(w[i] - 'a')
			if tmp.vals[0].first != x {
				// w肯定不是最大值的那部分
				res[j] = tr.vals[0].second
				break
			}
			res[j] = max(res[j], i+tmp.vals[1].second)
			// tmp.vals[0].first == x
			// no better way to choose the second path
			tmp = tmp.children[x]
			if tmp == nil || tmp.cnt == k {
				// 删除w后，不可能继续下去
				break
			}
			res[j] = max(res[j], i+1+tmp.vals[1].second)
			if i == len(w)-1 {
				res[j] = tr.vals[0].second
			}
		}
	}

	return res
}
