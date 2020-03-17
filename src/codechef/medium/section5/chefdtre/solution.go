package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, q := readTwoNums(scanner)

	solver := NewSolver(n)

	la := len("UNION")
	lb := len("GET")

	for q > 0 {
		q--
		scanner.Scan()
		bs := scanner.Bytes()
		var a, b, p int

		if bs[0] == 'U' {
			p = readInt(bs, la+1, &a)
			readInt(bs, p+1, &b)
			solver.Union(a, b)
		} else {
			p = readInt(bs, lb+1, &a)
			readInt(bs, p+1, &b)
			res := solver.Get(a, b)
			fmt.Println(res)
		}
	}
}

type Solver2 struct {
	nodes []*TreeNode
	idx   int
	n     int
}

func NewSolver2(n int) Solver2 {
	nodes := make([]*TreeNode, 3*n+3)

	for i := 1; i <= n; i++ {
		nodes[i] = NewTreeNode()
		nodes[i].Insert(1, n, i)
	}

	return Solver2{nodes, n + 1, n}
}

func (this *Solver2) Union(a, b int) {
	na := this.nodes[a]
	nb := this.nodes[b]
	this.nodes[this.idx] = na.Merge(nb)
	this.idx++
}

func (this *Solver2) Get(a, k int) int {
	return this.nodes[a].GetKth(k, 1, this.n)
}

type TreeNode struct {
	count       int
	left, right *TreeNode
}

func NewTreeNode() *TreeNode {
	node := new(TreeNode)
	return node
}

func (node *TreeNode) Insert(l int, r int, v int) {
	node.count++
	if l < r {
		mid := (l + r) >> 1
		if v <= mid {
			if node.left == nil {
				node.left = NewTreeNode()
			}
			node.left.Insert(l, mid, v)
		} else {
			if node.right == nil {
				node.right = NewTreeNode()
			}
			node.right.Insert(mid+1, r, v)
		}
	}
}

func (this *TreeNode) Merge(that *TreeNode) *TreeNode {
	this.count += that.count
	if that.left != nil {
		if this.left == nil {
			this.left = that.left
		} else {
			this.left = this.left.Merge(that.left)
		}
	}
	if that.right != nil {
		if this.right == nil {
			this.right = that.right
		} else {
			this.right = this.right.Merge(that.right)
		}
	}
	return this
}

func (this *TreeNode) GetKth(k int, l, r int) int {
	if l == r {
		return l
	}

	mid := (l + r) >> 1
	var count int
	if this.left != nil {
		count = this.left.count
	}
	if count >= k {
		return this.left.GetKth(k, l, mid)
	}

	return this.right.GetKth(k-count, mid+1, r)
}

type Solver struct {
	nodes       []*Node
	front, tail int
}

func NewSolver(N int) Solver {
	nodes := make([]*Node, 3*N+3)

	for i := 1; i <= N; i++ {
		node := NewNode(nil, i, nil)
		nodes[i] = node
	}

	return Solver{nodes, 1, N + 1}
}

func (solver *Solver) Union(a, b int) {
	sz := len(solver.nodes)
	a %= sz
	b %= sz
	c := solver.tail
	solver.nodes[c] = join2(solver.nodes[a], solver.nodes[b])
	solver.tail = (solver.tail + 1) % sz
}

func (solver *Solver) Get(a, k int) int {
	sz := len(solver.nodes)
	a %= sz

	node := solver.nodes[a].GetKth(k)

	return node.key
}

type Node struct {
	height      int
	key         int
	count       int
	left, right *Node
}

func (node *Node) GetHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *Node) GetCount() int {
	if node == nil {
		return 0
	}
	return node.count
}

func NewNode(l *Node, k int, r *Node) *Node {
	node := new(Node)
	node.left = l
	node.key = k
	node.right = r
	node.count = l.GetCount() + 1
	node.height = max(l.GetHeight(), r.GetHeight()) + 1
	return node
}

func (node *Node) GetKth(k int) *Node {
	if node == nil {
		return nil
	}
	if node.left.GetCount() == k-1 {
		return node
	}
	if k <= node.left.GetCount() {
		return node.left.GetKth(k)
	}

	return node.right.GetKth(k - node.GetCount())
}

func rotateRight(y *Node) *Node {
	x := y.left

	t2 := x.right

	x.right = y
	y.left = t2
	y.height = max(y.left.GetHeight(), y.right.GetHeight()) + 1

	y.count = y.left.GetCount() + 1

	x.height = max(x.left.GetHeight(), x.right.GetHeight()) + 1
	x.count = x.left.GetCount() + 1

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(x.left.GetHeight(), x.right.GetHeight()) + 1
	x.count = x.left.GetCount() + 1

	y.height = max(y.left.GetHeight(), y.right.GetHeight()) + 1
	y.count = y.GetCount() + 1

	return y
}

func joinRight(tl *Node, k int, tr *Node) *Node {
	l, kk, c := tl.left, tl.key, tl.right

	if c.GetHeight() <= tr.GetHeight()+1 {
		t := NewNode(c, k, tr)
		if t.GetHeight() <= l.GetHeight()+1 {
			return NewNode(l, kk, t)
		}
		return rotateLeft(NewNode(l, kk, rotateRight(t)))
	}
	tt := joinRight(c, k, tr)
	t := NewNode(l, kk, tt)

	if tt.GetHeight() <= l.GetHeight()+1 {
		return t
	}

	return rotateLeft(t)
}

func joinLeft(tl *Node, k int, tr *Node) *Node {
	c, kk, r := tr.left, tr.key, tl.right

	if c.GetHeight() <= tl.GetCount()+1 {
		t := NewNode(tl, k, c)
		if t.GetHeight() <= r.GetHeight()+1 {
			return NewNode(t, kk, r)
		}
		return rotateRight(NewNode(rotateLeft(t), kk, r))
	}

	tt := joinLeft(tl, k, c)
	t := NewNode(tt, kk, r)
	if tt.GetHeight() <= r.GetHeight()+1 {
		return t
	}
	return rotateRight(t)
}

func join(tl *Node, k int, tr *Node) *Node {
	if tl.GetHeight() > tr.GetHeight()+1 {
		return joinRight(tl, k, tr)
	}
	if tr.GetHeight() > tl.GetHeight()+1 {
		return joinLeft(tl, k, tr)
	}

	return NewNode(tl, k, tr)
}

func split(t *Node, k int) (*Node, bool, *Node) {
	if t == nil {
		return nil, false, nil
	}

	L, m, R := t.left, t.key, t.right

	if k == m {
		return L, true, R
	}
	if k < m {
		LL, b, RR := split(L, k)
		return LL, b, join(RR, m, R)
	}

	LL, b, RR := split(R, k)
	return join(LL, m, L), b, RR
}

func splitLast(t *Node) (*Node, int) {
	if t == nil {
		return nil, -1
	}
	L, k, R := t.left, t.key, t.right
	if R == nil {
		return L, k
	}
	TT, kk := splitLast(R)

	return join(L, k, TT), kk
}

func join2(tl, tr *Node) *Node {
	if tl == nil {
		return tr
	}

	LL, k := splitLast(tl)

	return join(LL, k, tr)
}

func union(t1, t2 *Node) *Node {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	tl, _, tr := split(t1, t2.key)
	nl := union(t1.left, tl)
	nr := union(t1.right, tr)

	return join(nl, t1.key, nr)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
