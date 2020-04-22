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

		solver := NewSolver(n, A)

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
