package p1044

import "sort"

const MOD = 1000000000000007

func longestDupSubstring1(S string) string {
	n := len(S)
	check := func(k int) int {
		pow := int64(1)
		hash := make(map[int64]bool)

		if k >= n {
			return -1
		}
		cur := int64(0)
		var i int
		for i < k {
			cur = (cur * 31) + int64(S[i]-'a')
			cur %= MOD
			pow = (pow * 31) % MOD
			i++
		}

		hash[cur] = true

		for i < n {
			cur = cur*31 + int64(S[i]-'a')
			cur = modSub(cur, (int64(S[i-k]-'a')*pow)%MOD)
			cur %= MOD
			if hash[cur] {
				return i
			}
			hash[cur] = true
			i++
		}
		return -1
	}

	k := sort.Search(n, func(k int) bool {
		return check(k) < 0
	})

	if k <= 1 {
		return ""
	}
	k--
	i := check(k)
	return S[i-k+1 : i+1]
}

func modSub(a, b int64) int64 {
	c := a - b
	if c < 0 {
		c += MOD
	}
	return c
}

func longestDupSubstring2(S string) string {
	S += "$"
	root := buildSuffixTree(S)

	var loop func(node *SuffixTreeNode, height int)

	var best int
	var pos int

	//n := len(S)

	loop = func(node *SuffixTreeNode, height int) {
		if node == nil {
			return
		}
		if node.index < 0 {
			// a leaf node
			for i := 0; i < len(node.children); i++ {
				if node.children[i] != nil {
					loop(node.children[i], height+node.children[i].EdgeLength())
				}
			}

		} else if best < height-node.EdgeLength() {
			best = height - node.EdgeLength()
			pos = node.index
		}
	}

	loop(root, 0)

	if best == 0 {
		return ""
	}
	return S[pos : pos+best]
}

type SuffixTreeNode struct {
	children   []*SuffixTreeNode
	suffixLink *SuffixTreeNode
	start      int
	end        *int
	index      int
}

func NewNode(start int, end *int) *SuffixTreeNode {
	node := new(SuffixTreeNode)
	node.children = make([]*SuffixTreeNode, 27)
	node.suffixLink = nil
	node.start = start
	node.end = end
	node.index = -1
	return node
}

func (node *SuffixTreeNode) EdgeLength() int {
	return *(node.end) - node.start + 1
}

func buildSuffixTree(text string) *SuffixTreeNode {
	//var root *SuffixTreeNode
	var lastNewNode *SuffixTreeNode

	// active point (node, edge, len)
	activeEdge := -1
	var activeLen int

	var reminder int
	var leafEnd int
	rootEnd := new(int)
	*rootEnd = -1

	root := NewNode(-1, rootEnd)
	activeNode := root

	walkDown := func(cur *SuffixTreeNode) bool {
		if activeLen >= cur.EdgeLength() {
			// move
			activeEdge += cur.EdgeLength()
			activeLen -= cur.EdgeLength()
			activeNode = cur
			return true
		}
		return false
	}

	createNode := func(start int, end *int) *SuffixTreeNode {
		node := NewNode(start, end)
		node.suffixLink = root
		return node
	}

	extend := func(pos int) {
		leafEnd = pos

		reminder++

		lastNewNode = nil

		for reminder > 0 {
			if activeLen == 0 {
				activeEdge = pos
			}
			j := letterIndex(text[activeEdge])
			if activeNode.children[j] == nil {
				activeNode.children[j] = createNode(pos, &leafEnd)

				if lastNewNode != nil {
					lastNewNode.suffixLink = activeNode
					lastNewNode = nil
				}
			} else {
				next := activeNode.children[j]
				if walkDown(next) {
					// start from next
					continue
				}
				if text[next.start+activeLen] == text[pos] {
					// rule 3
					if lastNewNode != nil && activeNode != root {
						lastNewNode.suffixLink = activeNode
						lastNewNode = nil
					}
					activeLen++
					break
				}

				// split
				splitEnd := next.start + activeLen - 1
				split := createNode(next.start, &splitEnd)
				activeNode.children[j] = split

				split.children[letterIndex(text[pos])] = createNode(pos, &leafEnd)
				next.start += activeLen

				split.children[letterIndex(text[next.start])] = next

				if lastNewNode != nil {
					lastNewNode.suffixLink = split
				}
				lastNewNode = split
			}

			reminder--
			if activeNode == root && activeLen > 0 {
				activeLen--
				activeEdge = pos - reminder + 1
			} else if activeNode != root {
				activeNode = activeNode.suffixLink
			}
		}
	}

	for i := 0; i < len(text); i++ {
		extend(i)
	}

	var setIndexDfs func(node *SuffixTreeNode, height int)

	setIndexDfs = func(node *SuffixTreeNode, height int) {
		if node == nil {
			return
		}

		leaf := true
		for i := 0; i < len(node.children); i++ {
			if node.children[i] != nil {
				leaf = false

				setIndexDfs(node.children[i], height+node.children[i].EdgeLength())
			}
		}

		if leaf {
			node.index = len(text) - height
		}
	}

	setIndexDfs(root, 0)

	return root
}

func letterIndex(b byte) int {
	if b == '$' {
		return 26
	}
	return int(b - 'a')
}

func getSuffixsFromTree(root *SuffixTreeNode, text string) []string {
	var res []string

	buf := make([]byte, len(text))

	var dfs func(pos int, node *SuffixTreeNode)

	n := len(text)

	dfs = func(pos int, node *SuffixTreeNode) {

		for i := node.start; i <= *node.end; i++ {
			buf[pos] = text[i]
			pos++
		}
		if *node.end == n-1 {
			res = append(res, string(buf[:pos]))
			return
		}

		for i := 0; i < len(node.children); i++ {
			if node.children[i] != nil {
				dfs(pos, node.children[i])
			}
		}
	}

	for i := 0; i < len(root.children); i++ {
		if root.children[i] != nil {
			dfs(0, root.children[i])
		}
	}

	return res
}

func longestDupSubstring(S string) string {
	suffixArray := buildSuffixArray(S)
	n := len(S)
	rs := make([]int, n)
	for i := 0; i < n; i++ {
		rs[suffixArray[i]] = i
	}

	lcp := make([]int, n)

	var k int
	for i := 0; i < n; i++ {
		if rs[i] == n-1 {
			k = 0
			continue
		}
		j := suffixArray[rs[i]+1]

		for i+k < n && j+k < n && S[i+k] == S[j+k] {
			k++
		}
		lcp[rs[i]] = k
		if k > 0 {
			k--
		}
	}

	best := 0
	pos := 0

	for i := 0; i < n; i++ {
		if lcp[i] > best {
			best = lcp[i]
			pos = i
		}
	}

	if best == 0 {
		return ""
	}
	return S[suffixArray[pos] : suffixArray[pos]+best]
}

func buildSuffixArray(str string) []int {
	n := len(str)

	ss := make(suffixs, n)
	for i := 0; i < n; i++ {
		ss[i] = NewSuffix(i)
		ss[i].rank[0] = int(str[i] - 'a')
		if i < n-1 {
			ss[i].rank[1] = int(str[i+1] - 'a')
		} else {
			ss[i].rank[1] = -1
		}
	}

	sort.Sort(ss)

	ii := make([]int, n)

	for k := 4; k < 2*n; k *= 2 {
		var rk int
		prev := ss[0].rank[0]
		ss[0].rank[0] = 0
		ii[ss[0].index] = 0

		for i := 1; i < n; i++ {
			if ss[i].rank[0] == prev && ss[i].rank[1] == ss[i-1].rank[1] {
				prev = ss[i].rank[0]
				ss[i].rank[0] = rk
			} else {
				prev = ss[i].rank[0]
				rk++
				ss[i].rank[0] = rk
			}

			ii[ss[i].index] = i
		}

		for i := 0; i < n; i++ {
			j := ss[i].index + k/2
			ss[i].rank[1] = -1
			if j < n {
				ss[i].rank[1] = ss[ii[j]].rank[0]
			}
		}
		sort.Sort(ss)
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = ss[i].index
	}
	return ans
}

type Suffix struct {
	index int
	rank  [2]int
}

func NewSuffix(index int) Suffix {
	return Suffix{index: index, rank: [...]int{0, 0}}
}

type suffixs []Suffix

func (this suffixs) Len() int {
	return len(this)
}

func (this suffixs) Less(i, j int) bool {
	return this[i].rank[0] < this[j].rank[0] || (this[i].rank[0] == this[j].rank[0] && this[i].rank[1] < this[j].rank[1])
}

func (this suffixs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
