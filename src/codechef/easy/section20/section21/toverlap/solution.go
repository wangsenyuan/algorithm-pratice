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
		n, m := readTwoNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, 2)
		}
		B := make([][]int, m)
		for i := 0; i < m; i++ {
			B[i] = readNNums(reader, 2)
		}
		res := solve2(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const MAX_X = 100000010

func solve(A [][]int, B [][]int) int64 {
	events := make(map[int]int)
	for _, a := range A {
		events[a[0]]++
		events[a[1]]++
	}

	for _, b := range B {
		events[b[0]]++
		events[b[1]]++
	}

	points := make([]int, 0, len(events))
	for k := range events {
		points = append(points, k)
	}

	sort.Ints(points)

	ii := make(map[int]int)

	for i := 0; i < len(points); i++ {
		ii[points[i]] = i
	}

	sumA := make([]int, len(points)+1)
	for _, a := range A {
		sumA[ii[a[0]]+1]++
		sumA[ii[a[1]]+1]--
	}
	sumB := make([]int, len(points)+1)
	for _, b := range B {
		sumB[ii[b[0]]+1]++
		sumB[ii[b[1]]+1]--
	}

	var a, b int
	var res int64

	for i := 1; i < len(points); i++ {
		a += sumA[i]
		b += sumB[i]
		res += int64(a) * int64(b) * int64(points[i]-points[i-1])
	}
	return res
}

func solve1(A [][]int, B [][]int) int64 {
	root := NewNode(0, MAX_X-1)

	for _, a := range A {
		u, v := a[0], a[1]
		u--
		v--
		root.update(u, v-1, 1)
	}

	var res int64

	for _, b := range B {
		u, v := b[0], b[1]
		u--
		v--
		res += root.get(u, v-1)
	}
	return res
}

func solve2(A [][]int, B [][]int) int64 {
	//root := NewSegTree(MAX_X)
	var x int
	for _, a := range A {
		x = max(x, a[1])
	}

	for _, b := range B {
		x = max(x, b[1])
	}

	root := NewSegTree(x + 1)

	for _, a := range A {
		u, v := a[0], a[1]
		u--
		v--
		root.update(u, v-1, 1)
	}

	var res int64

	for _, b := range B {
		u, v := b[0], b[1]
		u--
		v--
		res += root.get(u, v-1)
	}
	return res
}

type SegTree struct {
	n    int
	lazy []int
	val  []int64
}

func NewSegTree(n int) *SegTree {
	lazy := make([]int, 4*n)
	val := make([]int64, 4*n)

	return &SegTree{n, lazy, val}
}

func (tree *SegTree) push(i int, l, r int) {
	if tree.lazy[i] > 0 {
		if l < r {
			mid := (l + r) / 2
			j := i << 1
			tree.val[j] += int64(mid-l+1) * int64(tree.lazy[i])
			tree.val[j|1] += int64(r-mid) * int64(tree.lazy[i])
			tree.lazy[j] += tree.lazy[i]
			tree.lazy[j|1] += tree.lazy[i]
		}
		tree.lazy[i] = 0
	}
}

func (tree *SegTree) update(ll, rr, v int) {
	var loop func(i int, l, r int)

	loop = func(i, l, r int) {
		if rr < l || r < ll {
			return
		}
		if ll <= l && r <= rr {
			tree.val[i] += int64(v) * int64(r-l+1)
			tree.lazy[i] += v
			return
		}
		tree.push(i, l, r)
		mid := (l + r) / 2
		loop(i<<1, l, mid)
		loop((i<<1)|1, mid+1, r)
		tree.val[i] = tree.val[i<<1] + tree.val[(i<<1)|1]
	}
	loop(1, 0, tree.n-1)
}

func (tree *SegTree) get(ll, rr int) int64 {
	var loop func(i int, l, r int) int64
	loop = func(i int, l, r int) int64 {
		if rr < l || r < ll {
			return 0
		}
		if ll <= l && r <= rr {
			return tree.val[i]
		}
		tree.push(i, l, r)
		mid := (l + r) / 2
		a := loop(i<<1, l, mid)
		b := loop((i<<1)|1, mid+1, r)
		return a + b
	}
	return loop(1, 0, tree.n-1)
}

type Node struct {
	left, right *Node
	start, end  int
	lazy        int
	val         int64
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	return node
}

func (node *Node) incrementValue(v int) {
	node.val += int64(v) * int64(node.end-node.start+1)
}

func (node *Node) push() {
	if node == nil {
		return
	}
	if node.start < node.end {
		mid := (node.start + node.end) / 2
		if node.left == nil {
			node.left = NewNode(node.start, mid)
			node.right = NewNode(mid+1, node.end)
		}
		if node.lazy > 0 {
			node.left.lazy += node.lazy
			node.right.lazy += node.lazy
			node.left.incrementValue(node.lazy)
			node.right.incrementValue(node.lazy)
		}
	}
	node.lazy = 0
}

func (node *Node) getValue() int64 {
	if node == nil {
		return 0
	}
	return node.val
}

func (node *Node) update(ll int, rr int, v int) {
	if node.end < ll || rr < node.start || ll > rr {
		return
	}

	if ll <= node.start && node.end <= rr {
		node.val += int64(v) * int64(node.end-node.start+1)
		node.lazy += v
		return
	}
	node.push()

	node.left.update(max(ll, node.left.start), min(rr, node.left.end), v)
	node.right.update(max(ll, node.right.start), min(rr, node.right.end), v)

	node.val = node.left.getValue() + node.right.getValue()
}

func (node *Node) get(ll int, rr int) int64 {
	if node == nil || node.end < ll || rr < node.start || ll > rr {
		return 0
	}
	if ll <= node.start && node.end <= rr {
		return node.val
	}
	node.push()

	a := node.left.get(max(ll, node.left.start), min(rr, node.left.end))
	b := node.right.get(max(ll, node.right.start), min(rr, node.right.end))
	return a + b
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
