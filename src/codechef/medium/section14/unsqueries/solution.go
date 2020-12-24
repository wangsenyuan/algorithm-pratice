package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, q, s := readThreeNums(reader)
	heights := readNNums(reader, n)
	solver := NewSolver(n, s, heights)
	var buf bytes.Buffer
	for q > 0 {
		q--
		x, y := readTwoNums(reader)
		buf.WriteString(fmt.Sprintf("%d\n", solver.Query(x, y)))
	}
	fmt.Print(buf.String())
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

type Solver1 struct {
	n          int
	s          int
	last       int
	go_forward []int
	go_back    []int
	tree_gf    []int
	tree_gb    []int
	tree_hi    []int
	heights    []int
}

func NewSolver1(n int, s int, heights []int) *Solver1 {
	tmp := make([]int, n+1)
	copy(tmp[1:], heights)
	heights = tmp

	go_forward := make([]int, n+2)
	go_back := make([]int, n+1)
	stack := make([]int, n+1)
	next := make([]int, n+2)
	var p int
	for i := 1; i <= n; i++ {
		go_back[i] = 1
		next[i] = n + 1
		for p > 0 && heights[stack[p-1]] < heights[i] {
			next[stack[p-1]] = i
			go_back[i] = max(go_back[i], 1+go_back[stack[p-1]])
			p--
		}

		stack[p] = i
		p++
	}

	go_forward[n+1] = 0
	for i := n; i > 0; i-- {
		go_forward[i] = 1 + go_forward[next[i]]
	}
	tree_hi := make([]int, 2*n)
	buildHi(tree_hi, heights, n)

	tree_gf := make([]int, 2*n)
	build(tree_gf, go_forward, n)

	tree_gb := make([]int, 2*n)
	build(tree_gb, go_back, n)

	return &Solver1{n, s, 0, go_forward, go_back, tree_gf, tree_gb, tree_hi, heights}
}

func (this *Solver1) Query(x, y int) int {
	L := getPosition(x, this.s, this.last, this.n)
	R := getPosition(y, this.s, this.last, this.n)
	if L > R {
		L, R = R, L
	}
	hi := queryHi(this.tree_hi, this.heights, this.n, L-1, R)
	res := query(this.tree_gf, this.n, L-1, hi) - this.go_forward[hi] + 1
	if hi < R {
		res = max(res, query(this.tree_gb, this.n, hi, R))
	}
	this.last = res
	return res
}

func queryHi(t []int, H []int, n int, l int, r int) int {
	res := l

	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			if H[t[l]] >= H[res] {
				res = t[l]
			}
			l++
		}
		if r&1 == 1 {
			r--
			if H[t[r]] >= H[res] {
				res = t[r]
			}
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func buildHi(t []int, H []int, n int) {
	for i := 0; i < n; i++ {
		t[i+n] = i + 1
	}
	for i := n - 1; i > 0; i-- {
		left := i << 1
		right := left | 1
		if H[t[left]] >= H[t[right]] {
			t[i] = t[left]
		} else {
			t[i] = t[right]
		}
	}
}

func query(t []int, n int, l int, r int) int {
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func build(t []int, a []int, n int) {
	for i := 0; i < n; i++ {
		t[i+n] = a[i+1]
	}
	for i := n - 1; i > 0; i-- {
		t[i] = max(t[i<<1], t[i*2+1])
	}
}

type Solver struct {
	n     int
	s     int
	last  int
	nodes []*Node
}

type Node struct {
	left, right *Node
	val         int
	lazy        int
}

func copyNode(node *Node) *Node {
	if node == nil {
		return new(Node)
	}

	res := new(Node)
	res.left = node.left
	res.right = node.right
	res.val = node.val
	res.lazy = node.lazy
	return res
}

func (this *Node) push(l, r int) {
	this.left = copyNode(this.left)
	this.right = copyNode(this.right)
	if this.lazy > 0 {
		this.val += this.lazy
		if l < r {
			this.left.lazy += this.lazy
			this.right.lazy += this.lazy
		}
		this.lazy = 0
	}
}

func (node *Node) increment(l, r int, tl, tr int) {
	node.push(l, r)
	if l > r || tl > r || l > tr {
		return
	}
	if tl <= l && r <= tr {
		node.val++
		node.left.lazy++
		node.right.lazy++
		return
	}
	mid := (l + r) >> 1
	node.left.increment(l, mid, tl, tr)
	node.right.increment(mid+1, r, tl, tr)
	node.val = max(node.left.val, node.right.val)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (node *Node) get(l, r int, tl int, tr int) int {
	if node == nil {
		return 0
	}
	if l > r || tl > r || l > tr {
		return 0
	}

	node.push(l, r)

	if tl <= l && r <= tr {
		return node.val
	}
	mid := (l + r) >> 1
	a := node.left.get(l, mid, tl, tr)
	b := node.right.get(mid+1, r, tl, tr)
	return max(a, b)
}

func NewSolver(n int, s int, heights []int) *Solver {

	nodes := make([]*Node, n+1)
	nodes[0] = copyNode(nil)
	//ii := make([]int, n+1)
	stack := make([]int, n)
	var p int

	for r := 1; r <= n; r++ {
		for p > 0 && heights[stack[p-1]-1] < heights[r-1] {
			p--
		}
		nodes[r] = copyNode(nodes[r-1])

		var j int
		if p > 0 {
			j = stack[p-1]
		}
		nodes[r].increment(1, n, j+1, r)
		stack[p] = r
		p++
	}

	return &Solver{n, s, 0, nodes}
}

func getPosition(x int, s int, last int, n int) int {
	return (x+s*last-1)%n + 1
}

func (this *Solver) Query(x, y int) int {
	L := getPosition(x, this.s, this.last, this.n)
	R := getPosition(y, this.s, this.last, this.n)
	if L > R {
		L, R = R, L
	}
	tree := this.nodes[R]
	this.last = tree.get(1, this.n, L, R)
	return this.last
}
