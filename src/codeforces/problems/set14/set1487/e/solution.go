package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := readNNums(reader, 4)
	A := readNNums(reader, first[0])
	B := readNNums(reader, first[1])
	C := readNNums(reader, first[2])
	D := readNNums(reader, first[3])

	m := readNum(reader)
	X := make([][]int, m)
	for i := 0; i < m; i++ {
		X[i] = readNNums(reader, 2)
	}
	m = readNum(reader)
	Y := make([][]int, m)
	for i := 0; i < m; i++ {
		Y[i] = readNNums(reader, 2)
	}
	m = readNum(reader)
	Z := make([][]int, m)
	for i := 0; i < m; i++ {
		Z[i] = readNNums(reader, 2)
	}
	fmt.Println(solve1(A, B, C, D, X, Y, Z))
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

const MAX_N = 150010
const INF = 400000010

func solve(A, B, C, D []int, X, Y, Z [][]int) int {

	proc := func(a, b []int, c [][]int) []int {
		n1 := len(a)
		n2 := len(b)
		m := len(c)
		g := NewGraph(n1+n2, m)

		for _, x := range c {
			u, v := x[0], x[1]
			u--
			v--
			g.AddEdge(u, v+n1)
		}
		var set *Node
		for i := 0; i < n2; i++ {
			if b[i] < INF {
				set = Insert(set, b[i])
			}
		}
		res := make([]int, n1)
		for u := 0; u < n1; u++ {
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i] - n1
				if b[v] < INF {
					set = Delete(set, b[v])
				}
			}
			res[u] = INF
			if set != nil {
				res[u] = MinValueNode(set).key + a[u]
			}

			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i] - n1
				if b[v] < INF {
					set = Insert(set, b[v])
				}
			}
		}
		return res
	}

	R := proc(C, D, Z)
	R = proc(B, R, Y)
	R = proc(A, R, X)
	var best int = INF
	for i := 0; i < len(R); i++ {
		best = min(best, R[i])
	}
	if best >= INF {
		return -1
	}
	return best
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

/**
* this is a AVL tree
 */
type Node struct {
	key         int
	height      int
	cnt         int
	left, right *Node
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func NewNode(key int) *Node {
	node := new(Node)
	node.key = key
	node.height = 1
	node.cnt = 1
	return node
}

func rightRotate(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.Height(), y.right.Height()) + 1
	x.height = max(x.left.Height(), x.right.Height()) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.Height(), x.right.Height()) + 1
	y.height = max(y.left.Height(), y.right.Height()) + 1

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
		return node
	}

	if node.key > key {
		node.left = Insert(node.left, key)
	} else {
		node.right = Insert(node.right, key)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
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
		return root
	}

	root.height = max(root.left.Height(), root.right.Height()) + 1
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

func solve1(A, B, C, D []int, X, Y, Z [][]int) int {
	R := process(C, D, Z)
	R = process(B, R, Y)
	R = process(A, R, X)
	var best int = INF
	for i := 0; i < len(R); i++ {
		best = min(best, R[i])
	}
	if best >= INF {
		return -1
	}
	return best
}

func process(A, B []int, X [][]int) []int {
	n1 := len(A)
	n2 := len(B)
	sort.Slice(X, func(i, j int) bool {
		return X[i][1] < X[j][1]
	})

	tree := NewTree(n2)

	for i := 0; i < n2; i++ {
		tree.Update(i, B[i])
	}

	cnt := make([]int, n1)

	for _, x := range X {
		cnt[x[0]-1]++
	}

	conn := make([][]int, n1)

	for i := 0; i < n1; i++ {
		conn[i] = make([]int, 0, cnt[i])
	}

	for _, x := range X {
		u, v := x[0], x[1]
		u--
		v--
		conn[u] = append(conn[u], v)
	}
	res := make([]int, n1)
	for u := 0; u < n1; u++ {
		res[u] = INF
		var p int
		for _, v := range conn[u] {
			if p < v {
				tmp := tree.Query(p, v)
				res[u] = min(res[u], tmp+A[u])
			}
			p = v + 1
		}
		if p < n2 {
			res[u] = min(res[u], tree.Query(p, n2)+A[u])
		}
	}
	return res
}

type Tree struct {
	arr []int
	n   int
}

func NewTree(n int) *Tree {
	arr := make([]int, n*2)
	for i := 0; i < len(arr); i++ {
		arr[i] = INF
	}
	return &Tree{arr, n}
}

func (tree *Tree) Update(p int, v int) {
	p += tree.n
	if tree.arr[p] > v {
		tree.arr[p] = v
		for p > 1 {
			tree.arr[p>>1] = min(tree.arr[p], tree.arr[p^1])
			p >>= 1
		}
	}
}

func (tree *Tree) Query(l, r int) int {
	l += tree.n
	r += tree.n
	res := INF
	for l < r {
		if l&1 == 1 {
			res = min(res, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type SegTree struct {
	arr  []int
	lazy []int
	n    int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 4*n)
	lazy := make([]int, 4*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = INF + 1
	}
	return &SegTree{arr, lazy, n}
}

func (tree *SegTree) Update(L, R int, V int) {
	var loop func(i int, l, r int)
	loop = func(i int, l, r int) {
		if r < L || R < l {
			return
		}
		if L <= l && r <= R {
			tree.change(i, V)
			return
		}
		tree.push(i)
		mid := (l + r) / 2
		loop(2*i, l, mid)
		loop(2*i+1, mid+1, r)
		tree.arr[i] = min(tree.arr[2*i], tree.arr[2*i+1])
	}
	loop(1, 0, tree.n-1)
}

func (tree *SegTree) Query(L, R int) int {
	var loop func(i int, l, r int) int
	loop = func(i int, l, r int) int {
		if r < L || R < l {
			return INF
		}
		if L <= l && r <= R {
			return tree.arr[i]
		}

		tree.push(i)

		mid := (l + r) / 2
		a := loop(2*i, l, mid)
		b := loop(2*i+1, mid+1, r)
		return min(a, b)
	}
	return loop(1, 0, tree.n-1)
}

func (tree *SegTree) change(id int, v int) {
	if tree.arr[id] > v {
		tree.arr[id] = v
		tree.lazy[id] = v
	}
}

func (tree *SegTree) push(id int) {
	if tree.lazy[id] > 0 {
		tree.change(id*2, tree.lazy[id])
		tree.change(id*2+1, tree.lazy[id])
		tree.lazy[id] = 0
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
