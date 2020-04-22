package main

import (
	"bufio"
	"fmt"
	"math"
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)

		solver := NewSolver1(n, A)

		for m > 0 {
			m--
			bs, _ := reader.ReadBytes('\n')
			var a, b int
			pos := 2
			pos = readInt(bs, pos, &a)
			readInt(bs, pos+1, &b)
			if bs[0] == 'Q' {
				fmt.Println(solver.Query(a, b))
			} else {
				solver.Update(a, b)
			}
		}
	}
}

type Solver1 struct {
	n  int
	A  []int
	sr *SegTree
	sl *SegTree
}

func NewSolver1(n int, A []int) Solver1 {
	sr := NewSegTree(n)
	sl := NewSegTree(n)

	for i := 0; i < n; i++ {
		sr.Update(i+1, n, int64(A[i]))
		sl.Update(i+1, n, -int64(A[i]))
	}

	return Solver1{n, A, sr, sl}
}

func (solver *Solver1) Update(pos int, v int) {
	solver.sr.Update(pos, solver.n, int64(v-solver.A[pos-1]))
	solver.sl.Update(pos, solver.n, int64(solver.A[pos-1]-v))
	solver.A[pos-1] = v
}

func (solver *Solver1) Query(left, right int) int64 {
	return solver.sr.Query(right, solver.n) + solver.sl.Query(0, left-1)
}

type SegTree struct {
	n int
	A []int64
	B []int64
}

func NewSegTree(n int) *SegTree {
	A := make([]int64, 4*n)
	B := make([]int64, 4*n)
	return &SegTree{n, A, B}
}

func (seg *SegTree) apply(i int, v int64) {
	seg.A[i] += v
	seg.B[i] += v
}

func (seg *SegTree) push(i int) {
	seg.apply(2*i, seg.B[i])
	seg.apply(2*i+1, seg.B[i])
	seg.B[i] = 0
}

func (seg *SegTree) Update(left, right int, v int64) {

	var loop func(i int, start int, end int)

	loop = func(i int, start int, end int) {
		if left <= start && end <= right {
			seg.apply(i, v)
			return
		}

		seg.push(i)
		mid := (start + end) / 2

		if left <= mid {
			loop(2*i, start, mid)
		}

		if mid < right {
			loop(2*i+1, mid+1, end)
		}

		seg.A[i] = max(seg.A[2*i], seg.A[2*i+1])
	}

	loop(1, 0, seg.n)
}

func (seg *SegTree) Query(left, right int) int64 {

	var loop func(i int, start, end int) int64

	loop = func(i int, start, end int) int64 {
		if left <= start && end <= right {
			return seg.A[i]
		}
		mid := (start + end) / 2
		seg.push(i)

		var x int64 = math.MinInt64

		if left <= mid {
			x = loop(2*i, start, mid)
		}

		var y int64 = math.MinInt64

		if mid < right {
			y = loop(2*i+1, mid+1, end)
		}

		return max(x, y)
	}

	return loop(1, 0, seg.n)
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Solver struct {
	n     int
	A     []int
	sum   *BIT
	left  *Node
	right *Node
}

func NewSolver(n int, A []int) Solver {
	sum := &BIT{n, make([]int64, n+1)}

	for i := 0; i < n; i++ {
		sum.Update(i, A[i])
	}

	left := NewNode(0, n-1)

	for i := 0; i < n; i++ {
		left.Update(i, int64(A[i]))
	}

	right := NewNode(0, n-1)

	for i := n - 1; i >= 0; i-- {
		right.Update(n-1-i, int64(A[i]))
	}

	return Solver{n, A, sum, left, right}
}

func (solver Solver) Query(left, right int) int64 {
	left--
	right--

	total := solver.sum.Get(solver.n - 1)

	var lSum int64
	if left > 0 {
		lSum = solver.sum.Get(left - 1)
	}
	rSum := solver.sum.Get(right)

	res := rSum - lSum

	if left > 0 {
		lMSum := solver.left.GetValue(left - 1)
		res += lSum - lMSum
	}

	rSum = total - rSum

	if right < solver.n-1 {
		rMSum := solver.right.GetValue(solver.n - 1 - (right + 1))
		res += rSum - rMSum
	}

	return res
}

func (solver *Solver) Update(pos int, v int) {
	V := int64(v)
	pos--

	solver.sum.Update(pos, v-solver.A[pos])

	solver.left.Update(pos, V)
	solver.right.Update(solver.n-1-pos, V)

	solver.A[pos] = v
}

type BIT struct {
	n   int
	arr []int64
}

func (bit *BIT) Update(pos int, v int) {
	pos++

	V := int64(v)

	for pos <= bit.n {
		bit.arr[pos] += V
		pos += pos & (-pos)
	}
}

func (bit *BIT) Get(pos int) int64 {
	pos++

	var res int64

	for pos > 0 {
		res += bit.arr[pos]
		pos -= pos & (-pos)
	}

	return res
}

type Node struct {
	start, end  int
	left, right *Node
	value       int64
	sum         int64
}

const INF = 1 << 30

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.value = INF
	node.sum = 0

	return node
}

func (node Node) GetMid() int {
	return (node.start + node.end) / 2
}

func (node *Node) push() {
	if node.start == node.end || node.left != nil {
		return
	}
	mid := node.GetMid()
	node.left = NewNode(node.start, mid)
	node.right = NewNode(mid+1, node.end)
}

func (node *Node) Update(pos int, value int64) {
	if node.start == node.end {
		node.value = min(0, value)
		node.sum = value
		return
	}
	node.push()

	mid := node.GetMid()

	if pos <= mid {
		node.left.Update(pos, value)
	} else {
		node.right.Update(pos, value)
	}
	node.value = min(node.left.value, node.left.sum+node.right.value)
	node.sum = node.left.sum + node.right.sum
}

func (node *Node) GetValue(pos int) int64 {
	// the value is actually [0...pos]
	if node.start > pos {
		return INF
	}

	if node.end <= pos {
		return node.value
	}

	// start <= pos && pos <= end
	mid := node.GetMid()

	if pos <= mid {
		return node.left.GetValue(pos)
	}

	tmp := node.right.GetValue(pos)
	res := min(node.left.value, node.left.sum+tmp)
	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
