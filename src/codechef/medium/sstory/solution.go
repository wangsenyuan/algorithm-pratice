package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := scanner.Bytes()
	scanner.Scan()
	s2 := scanner.Bytes()
	length, start := solve(s1, s2)
	if length == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(string(s2[start : start+length]))
		fmt.Println(length)
	}
}

func solve(s1, s2 []byte) (length, start int) {
	n1 := len(s1)
	st := make([]*State, 2*n1+1)
	st[0] = &State{0, -1, make([]int, 26)}
	last, size := 0, 1
	for i := 0; i < n1; i++ {
		c := int(s1[i] - 'a')
		cur := size
		st[cur] = &State{st[last].length + 1, 0, make([]int, 26)}
		size++
		p := last
		for p >= 0 && st[p].next[c] == 0 {
			st[p].next[c] = cur
			p = st[p].link
		}

		if p >= 0 {
			q := st[p].next[c]
			if st[p].length+1 == st[q].length {
				st[cur].link = q
			} else {
				r := size
				st[r] = &State{st[p].length + 1, st[q].link, make([]int, 26)}
				size++
				copy(st[r].next, st[q].next)
				for p >= 0 && st[p].next[c] == q {
					st[p].next[c] = r
					p = st[p].link
				}
				st[cur].link = r
				st[q].link = r
			}
		}
		last = cur
	}

	var v, cur int
	for i := 0; i < len(s2); i++ {
		x := int(s2[i] - 'a')
		for v > 0 && st[v].next[x] == 0 {
			v = st[v].link
			cur = st[v].length
		}
		if st[v].next[x] > 0 {
			v = st[v].next[x]
			cur++
		}
		if cur > length {
			length = cur
			start = i - cur + 1
		}
	}
	return
}

type State struct {
	length int
	link   int
	next   []int
}

func solve1(s1, s2 []byte) (length, start int) {
	var mod int64 = 34369934
	var magic int64 = 9584612342

	check := func(l int) int {
		if l == 0 {
			return -1
		}

		var hash int64
		var mul int64 = 1
		for i := 0; i < l; i++ {
			hash = (hash*magic + int64(s1[i])) % mod
			mul = (mul * magic) % mod
		}

		mul = (mod - mul) % mod
		hs := make(map[int64]bool)
		hs[hash] = true

		for i := l; i < len(s1); i++ {
			hash = (hash*magic + int64(s1[i]) + mul*int64(s1[i-l])) % mod
			hs[hash] = true
		}
		hash = 0

		for i := 0; i < l; i++ {
			hash = (hash*magic + int64(s2[i])) % mod
		}
		if hs[hash] {
			return 0
		}
		for i := l; i < len(s2); i++ {
			hash = (hash*magic + int64(s2[i]) + mul*int64(s2[i-l])) % mod
			if hs[hash] {
				return i - l
			}
		}
		return -1
	}

	left, right := 0, len(s2)

	for left < right {
		mid := (left + right) / 2
		res := check(mid)
		if res >= 0 {
			left = mid
		} else {
			right = mid - 1
		}
	}

	start = check(left)

	if start < 0 {
		return
	}

	length = left
	return
}

func solve2(s1, s2 []byte) (length int, start int) {
	items := make(Items, len(s1)+len(s2))

	for i := 0; i < len(s1)+len(s2); i++ {
		if i < len(s1) {
			items[i] = Item{s1, i, 1}
		} else {
			items[i] = Item{s2, i - len(s1), 2}
		}
	}

	sort.Sort(items)

	for i := len(items) - 2; i >= 0; i-- {
		a, b := items[i+1], items[i]
		if a.src == b.src {
			// same string
			continue
		}
		var k int
		for a.idx+k < len(a.s) && b.idx+k < len(b.s) && a.s[a.idx+k] == b.s[b.idx+k] {
			k++
		}

		c := a
		if b.src == 2 {
			c = b
		}

		if k > length || (k == length && c.idx < start) {
			length = k
			start = c.idx
		}
	}

	return
}

type Item struct {
	s   []byte
	idx int
	src int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	a, b := items[i].s, items[j].s
	ai, bi := items[i].idx, items[j].idx

	for ai < len(a) && bi < len(b) {
		if a[ai] < b[bi] {
			return true
		} else if a[ai] > b[bi] {
			return false
		}
		ai++
		bi++
	}

	if bi < len(b) {
		return true
	}
	return false
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func solve3(s1, s2 []byte) string {
	s3 := make([]byte, len(s1)+len(s2)+2)
	copy(s3, s2)
	s3[len(s2)] = '$'
	copy(s3[len(s2)+1:], s1)
	s3[len(s1)+1+len(s2)] = '#'
	st := BuildSuffixTree(string(s3))
	return LCS(st, len(s2))
}

type SuffixTreeNode struct {
	children    []*SuffixTreeNode
	link        *SuffixTreeNode
	start       int
	end         *int
	suffixIndex int
}

type SuffixTree struct {
	root                 *SuffixTreeNode
	active               *SuffixTreeNode
	activeEdge           int
	activeLength         int
	remainingSuffixCount int
	leafEnd              int
	str                  string
}

func (st *SuffixTree) newNode(start int, end *int) *SuffixTreeNode {
	node := &SuffixTreeNode{}
	// 26 + 2
	node.children = make([]*SuffixTreeNode, 28)
	node.link = st.root
	node.end = end
	node.start = start
	node.suffixIndex = -1
	return node
}

func (stn SuffixTreeNode) edgeLength() int {
	return *(stn.end) - stn.start + 1;
}

func (st *SuffixTree) walkDown(current *SuffixTreeNode) bool {
	if st.activeLength >= current.edgeLength() {
		st.activeEdge += current.edgeLength()
		st.activeLength -= current.edgeLength()
		st.active = current
		return true
	}
	return false
}

func indexOfLetter(b byte) int {
	if b == '$' {
		return 26
	} else if b == '#' {
		return 27
	}
	return int(b - 'a')
}

func (st *SuffixTree) extendSuffixTree(pos int) {
	st.leafEnd = pos
	st.remainingSuffixCount++
	var last *SuffixTreeNode

	for st.remainingSuffixCount > 0 {
		if st.activeLength == 0 {
			st.activeEdge = pos
		}
		if st.active.children[indexOfLetter(st.str[st.activeEdge])] == nil {
			st.active.children[indexOfLetter(st.str[st.activeEdge])] = st.newNode(pos, &st.leafEnd)
			if last != nil {
				last.link = st.active
				last = nil
			}
		} else {
			next := st.active.children[indexOfLetter(st.str[st.activeEdge])]
			if st.walkDown(next) {
				continue
			}
			if st.str[next.start+st.activeLength] == st.str[pos] {
				if last != nil && st.active != st.root {
					last.link = st.active
					last = nil
				}
				st.activeLength++
				break
			}
			splitEnd := next.start + st.activeLength - 1
			split := st.newNode(next.start, &splitEnd)
			st.active.children[indexOfLetter(st.str[st.activeEdge])] = split

			split.children[indexOfLetter(st.str[pos])] = st.newNode(pos, &st.leafEnd)
			next.start += st.activeLength
			split.children[indexOfLetter(st.str[pos])] = next

			if last != nil {
				last.link = split
			}
			last = split
		}

		st.remainingSuffixCount--

		if st.active == st.root && st.activeLength > 0 {
			st.activeLength--
			st.activeEdge = pos - st.remainingSuffixCount + 1
		} else if st.active != st.root {
			st.active = st.active.link
		}
	}
}

func (st *SuffixTree) setSuffixIndex() {
	var process func(node *SuffixTreeNode, labelHeight int)

	process = func(node *SuffixTreeNode, labelHeight int) {
		if node == nil {
			return
		}
		leaf := true
		for i := 0; i < 28; i++ {
			if node.children[i] != nil {
				leaf = false
				process(node.children[i], labelHeight+node.children[i].edgeLength())
			}
		}
		if leaf {
			for i := node.start; i <= *(node.end); i++ {
				if st.str[i] == '$' {
					j := i
					node.end = &j
				}
			}
			node.suffixIndex = len(st.str) - labelHeight
		}
	}

	process(st.root, 0)
}

func BuildSuffixTree(str string) *SuffixTree {
	st := &SuffixTree{}
	rootEnd := -1
	st.root = st.newNode(-1, &rootEnd)
	st.active = st.root
	st.str = str
	for i := 0; i < len(str); i++ {
		st.extendSuffixTree(i)
	}
	st.setSuffixIndex()
	return st
}

func LCS(st *SuffixTree, size1 int) string {
	var dfs func(node *SuffixTreeNode, labelHeight int) int
	var maxHeight, start int

	dfs = func(node *SuffixTreeNode, labelHeight int) int {
		if node == nil {
			return 0
		}

		if node.suffixIndex < 0 {
			for i := 0; i < 28; i++ {
				if node.children[i] != nil {
					ret := dfs(node.children[i], labelHeight+node.children[i].edgeLength())
					if node.suffixIndex == -1 {
						node.suffixIndex = ret
					} else if (node.suffixIndex == -2 && ret == -3) || (node.suffixIndex == -3 && ret == -2) || node.suffixIndex == -4 {
						node.suffixIndex = -4
						nodeStart := *(node.end) - labelHeight + 1
						if maxHeight < labelHeight {
							maxHeight = labelHeight
							start = nodeStart
						} else if maxHeight == labelHeight && nodeStart < start {
							maxHeight = labelHeight
							start = nodeStart
						}
					}
				}
			}
		}

		if node.suffixIndex > -1 && node.suffixIndex < size1 {
			return -2
		} else if node.suffixIndex > -1 && node.suffixIndex >= size1 {
			return -3
		}
		return node.suffixIndex
	}

	dfs(st.root, 0)

	if maxHeight > 0 {
		return st.str[start : start+maxHeight]
	}
	return ""
}
