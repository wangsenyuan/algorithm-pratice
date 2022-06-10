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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		S := make([][]int, n)
		for i := 0; i < n; i++ {
			S[i] = readNNums(reader, 3)
		}
		res := solve(n, S)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve1(n int, S [][]int) int64 {
	segs := make([]Seg, n)
	for i := 0; i < n; i++ {
		segs[i] = Seg{S[i][0], S[i][1], S[i][2]}
	}
	sort.Sort(Segs2(segs))

	arr := make([]Pair, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = Pair{0, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i*2].Min(arr[i*2+1])
	}

	query := func(l, r int) Pair {
		l += n
		r += n
		ans := arr[l]

		for l < r {
			if l&1 == 1 {
				ans = ans.Min(arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				ans = ans.Min(arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return ans
	}

	set := func(p int, v Pair) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = arr[p].Min(arr[p^1])
			p >>= 1
		}
	}

	for _, seg := range segs {
		l, r := seg.l, seg.r
		l--
		x := query(l, r)
		if x.first < seg.v {
			set(x.second, Pair{seg.v, x.second})
		}
	}
	var res int64

	for i := 0; i < n; i++ {
		x := query(i, i+1)
		res += int64(x.first)
	}
	return res
}

type Pair struct {
	first, second int
}

func (this Pair) Min(that Pair) Pair {
	if this.first < that.first {
		return this
	}
	return that
}

func solve(n int, S [][]int) int64 {
	segs := make([]Seg, n+1)
	for i := 0; i < n; i++ {
		segs[i] = Seg{S[i][0], S[i][1], S[i][2]}
	}
	segs[n] = Seg{1, n, 0}
	sort.Sort(Segs(segs))

	within := func(a Seg, b Seg) bool {
		return a.l <= b.l && b.r <= a.r
	}

	n++

	g := NewGraph(n, n)

	var dfs func(p int)

	var i int

	dfs = func(p int) {
		if i == n {
			return
		}
		if p != i {
			g.AddEdge(p, i)
		}

		i++

		for i < n && within(segs[i-1], segs[i]) {
			dfs(i - 1)
		}
		for i < n && within(segs[p], segs[i]) {
			dfs(p)
		}
		// not a child
	}

	dfs(0)

	var dfs2 func(u int)

	roots := make([]*Node, n)

	dfs2 = func(u int) {
		big := -1
		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dfs2(v)
			if big < 0 || roots[big].Size() < roots[v].Size() {
				big = v
			}
		}

		if big >= 0 {
			roots[u] = roots[big]
		}

		roots[u] = Insert(roots[u], segs[u].v)

		for i := g.node[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v != big {
				Iterate(roots[v], func(x int) {
					roots[u] = Insert(roots[u], x)
				})
			}
		}

		for roots[u].Size() > segs[u].Count() {
			// already full
			minValue := roots[u].MinValue()
			roots[u] = Delete(roots[u], minValue)
		}
	}

	var res int64

	dfs2(0)

	Iterate(roots[0], func(x int) {
		res += int64(x)
	})
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Graph struct {
	node []int
	next []int
	to   []int
	cur  int
}

func NewGraph(n int, e int) *Graph {
	g := new(Graph)
	g.node = make([]int, n)
	g.next = make([]int, e+1)
	g.to = make([]int, e+1)
	return g
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.node[u]
	g.node[u] = g.cur
	g.to[g.cur] = v
}

type Seg struct {
	l int
	r int
	v int
}

func (seg Seg) Count() int {
	return seg.r - seg.l + 1
}

type Segs2 []Seg

func (segs Segs2) Len() int {
	return len(segs)
}

func (segs Segs2) Less(i, j int) bool {
	return segs[i].Count() < segs[j].Count()
}

func (segs Segs2) Swap(i, j int) {
	segs[i], segs[j] = segs[j], segs[i]
}

type Segs []Seg

func (segs Segs) Len() int {
	return len(segs)
}

func (segs Segs) Less(i, j int) bool {
	return segs[i].l < segs[j].l ||
		segs[i].l == segs[j].l && segs[i].r > segs[j].r
}

func (segs Segs) Swap(i, j int) {
	segs[i], segs[j] = segs[j], segs[i]
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	size        int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	node.size = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + y.cnt

	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + x.cnt

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	x.size = x.left.Size() + x.right.Size() + x.cnt

	y.height = max(y.left.Height(), y.right.Height()) + 1
	y.size = y.left.Size() + y.right.Size() + y.cnt

	return y
}

func (node *Node) GetBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func Insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		node.cnt++
		node.size++
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.size = node.left.Size() + node.right.Size() + node.cnt

	balance := node.GetBalance()

	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func MinValueNode(root *Node) *Node {
	cur := root

	for cur.left != nil {
		cur = cur.left
	}

	return cur
}

func (root *Node) MinValue() int {
	if root == nil {
		return 0
	}
	node := MinValueNode(root)

	return node.key
}

func Delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = Delete(root.left, key)
	} else if key > root.key {
		root.right = Delete(root.right, key)
	} else {
		root.cnt--
		root.size--
		if root.cnt > 0 {
			return root
		}
		if root.left == nil || root.right == nil {
			tmp := root.left
			if root.left == nil {
				tmp = root.right
			}
			root = tmp
		} else {
			tmp := MinValueNode(root.right)

			root.key = tmp.key
			root.cnt = tmp.cnt
			// make sure tmp node deleted after call delete on root.right
			tmp.cnt = 1
			root.right = Delete(root.right, tmp.key)
		}
	}

	if root == nil {
		return nil
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
	root.size = root.left.Size() + root.right.Size() + root.cnt
	balance := root.GetBalance()

	if balance > 1 && root.left.GetBalance() >= 0 {
		return rightRotate(root)
	}

	if balance > 1 && root.left.GetBalance() < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	if balance < -1 && root.right.GetBalance() <= 0 {
		return leftRotate(root)
	}

	if balance < -1 && root.right.GetBalance() > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

func Iterate(node *Node, fn func(x int)) {
	if node == nil {
		return
	}
	for i := 0; i < node.cnt; i++ {
		fn(node.key)
	}
	Iterate(node.left, fn)
	Iterate(node.right, fn)
}
