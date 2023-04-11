package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 5)
		}
		res := solve(A, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A []int, Q [][]int) []int {
	// 3-d segment-tree
	values := make([]int, len(A))
	copy(values, A)
	values = sortUnique(values)
	index := createIndexer(values)

	for i := 0; i < len(A); i++ {
		A[i] = index[A[i]]
	}

	n := len(values)

	prev := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = -1
	}

	trees := BuildTrees(0, n)

	for i, a := range A {
		trees.AddPoint(a, prev[a], i)
		prev[a] = i
	}

	trees.seal()

	var ans int

	var res []int

	N := int64(len(A))

	for _, q := range Q {
		a, b, c, d, k := q[0], q[1], q[2], q[3], q[4]
		l := int((int64(max(ans, 0))*int64(a) + int64(b)) % N)
		r := int((int64(max(ans, 0))*int64(c)+int64(d))%N) + 1
		L, R := 0, n
		cur := trees
		if cur.Count(l, r) < k {
			ans = -1
		} else {
			for R-L > 1 {
				mid := (L + R) / 2
				ct := cur.left.Count(l, r)
				if k > ct {
					k -= ct
					L = mid
					cur = cur.right
				} else {
					R = mid
					cur = cur.left
				}
			}
			ans = L
		}
		if ans >= 0 {
			ans = values[ans]
		}
		res = append(res, ans)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// leaf node
type Node struct {
	left, right *Node
	i           int
	j           int
	c           int
}

func NewNode(i int, j int, c int) *Node {
	n := new(Node)
	n.i = i
	n.j = j
	n.c = c
	return n
}

func combine(left, right *Node) *Node {
	node := NewNode(left.i, right.j, left.c+right.c)
	node.left = left
	node.right = right
	return node
}

func (node *Node) AddValue(v int) *Node {
	if node.i <= v && v < node.j {
		if node.left != nil {
			return combine(node.left.AddValue(v), node.right.AddValue(v))
		}
		return NewNode(node.i, node.j, node.c+1)
	}
	return node
}

func (node *Node) Query(v int) int {
	if node.j <= v {
		return node.c
	}
	if v <= node.i {
		return 0
	}
	return node.left.Query(v) + node.right.Query(v)
}

func BuildSegTree(i int, j int) *Node {
	if j-i == 1 {
		return NewNode(i, j, 0)
	}
	k := (i + j) >> 1
	return combine(BuildSegTree(i, k), BuildSegTree(k, j))
}

func sortUnique(arr []int) []int {
	sort.Ints(arr)
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}

func createIndexer(arr []int) map[int]int {
	res := make(map[int]int)
	for i, x := range arr {
		res[x] = i
	}
	return res
}

func find(arr []int, v int) int {
	l, r := -1, len(arr)

	for r-l > 1 {
		mid := (l + r) >> 1
		if arr[mid] < v {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

type Pair struct {
	first  int
	second int
}

type Tree struct {
	xs     []int
	ys     []int
	points []Pair
	nodes  []*Node
}

func NewTree() *Tree {
	t := new(Tree)
	t.xs = make([]int, 0, 1)
	t.ys = make([]int, 0, 1)
	t.points = make([]Pair, 0, 1)
	t.nodes = make([]*Node, 0, 1)
	return t
}

func (t *Tree) addPoint(x, y int) {
	t.points = append(t.points, Pair{x, y})
	t.xs = append(t.xs, x)
	t.ys = append(t.ys, y)
}

func (t *Tree) seal() {
	// stop further updates
	sort.Slice(t.points, func(i, j int) bool {
		return t.points[i].first < t.points[j].first ||
			t.points[i].first == t.points[j].first && t.points[i].second < t.points[j].second
	})
	t.xs = sortUnique(t.xs)
	t.ys = sortUnique(t.ys)
	xInd := createIndexer(t.xs)
	yInd := createIndexer(t.ys)
	t.nodes = make([]*Node, len(t.xs))
	cur := BuildSegTree(0, len(t.ys))
	for i := 0; i < len(t.points); i++ {
		x := t.points[i].first
		y := t.points[i].second
		t.nodes[xInd[x]] = cur.AddValue(yInd[y])
		cur = t.nodes[xInd[x]]
	}
}

func (t *Tree) Count(l int, r int) int {
	i := find(t.xs, l) - 1
	l = find(t.ys, l)
	r = find(t.ys, r)
	return t.nodes[i].Query(r) - t.nodes[i].Query(l)
}

// 3d tree
type Trees struct {
	left, right *Trees
	l, r        int
	tree        *Tree
}

func BuildTrees(l int, r int) *Trees {
	t := new(Trees)
	t.l = l
	t.r = r
	t.tree = NewTree()

	if r-l == 1 {
		return t
	}
	mid := (l + r) >> 1
	t.left = BuildTrees(l, mid)
	t.right = BuildTrees(mid, r)
	return t
}

func (t *Trees) AddPoint(x int, y int, z int) {
	if t.l <= x && x < t.r {
		t.tree.addPoint(y, z)
		if t.left != nil {
			t.left.AddPoint(x, y, z)
			t.right.AddPoint(x, y, z)
		}
	}
}

func (t *Trees) seal() {
	t.tree.seal()
	if t.left != nil {
		t.left.seal()
		t.right.seal()
	}
}

func (t *Trees) Count(l int, r int) int {
	return t.tree.Count(l, r)
}
