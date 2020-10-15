package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)

		A := readNNums(reader, n)

		solver := NewSolver(n, A)

		for q > 0 {
			q--
			s, _ := reader.ReadBytes('\n')
			var id, pos int
			pos = readInt(s, 2, &id)
			if s[0] == '?' {
				var L, R int
				pos = readInt(s, pos+1, &L)
				readInt(s, pos+1, &R)
				ans := solver.GetAnswer(id, L, R)
				buf.WriteString(fmt.Sprintf("%d\n", ans))
			} else {
				var x int
				readInt(s, pos+1, &x)
				solver.Add(id, x)
			}
		}
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

type Solver struct {
	root *Node
	N    int
}

func NewSolver(n int, A []int) Solver {

	var build func(start int, end int) *Node

	build = func(start int, end int) *Node {
		if start == end-1 {
			node := NewNode(start, end)
			node.mx = A[start]
			return node
		}
		node := NewNode(start, end)
		mid := (start + end) / 2
		node.left = build(start, mid)
		node.right = build(mid, end)
		node.mx = max(node.left.mx, node.right.mx)
		node.ans = node.left.ans + count(node.right, node.left.mx)
		return node
	}

	root := build(0, n)

	return Solver{root, n}
}

func (node *Node) Add(pos int, x int) {
	if node.end <= pos || pos < node.start {
		return
	}
	if node.start == node.end-1 {
		node.mx += x
		return
	}

	mid := (node.start + node.end) / 2

	if pos < mid {
		node.left.Add(pos, x)
	} else {
		node.right.Add(pos, x)
	}

	node.mx = max(node.left.mx, node.right.mx)
	node.ans = node.left.ans + count(node.right, node.left.mx)
}

func (solver *Solver) Add(i int, x int) {
	i--
	solver.root.Add(i, x)
}

func (node *Node) GetRBound(l int, R int) int {
	if node.mx < R {
		return node.end
	}
	if node.start == node.end-1 {
		return node.start
	}

	if node.left.end <= l {
		return node.right.GetRBound(l, R)
	}
	ret := node.left.GetRBound(l, R)

	if ret < node.left.end {
		return ret
	}
	return node.right.GetRBound(l, R)
}

func (solver Solver) GetAnswer(id int, L int, R int) int {
	id--
	h := solver.root.GetRBound(id, R)
	ans := solver.root.Get(id, min(solver.N, h+1), L-1)

	return ans.first
}

func (node *Node) Get(zac int, kon int, maxFromLeft int) Pair {
	if zac >= node.end || kon <= node.start {
		return Pair{0, -1}
	}
	if zac == node.start && kon == node.end {
		return Pair{count(node, maxFromLeft), node.mx}
	}

	a := node.left.Get(zac, min(kon, node.left.end), maxFromLeft)
	b := node.right.Get(max(zac, node.right.start), kon, max(maxFromLeft, a.second))

	return Pair{a.first + b.first, max(a.second, b.second)}
}

func count(node *Node, mxFromLeft int) int {
	if node.start == node.end-1 {
		if node.mx > mxFromLeft {
			return 1
		}
		return 0
	}

	if node.left.mx <= mxFromLeft {
		return count(node.right, mxFromLeft)
	}
	ret := count(node.left, mxFromLeft)
	ret += node.ans - node.left.ans
	return ret
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

type Node struct {
	left, right *Node
	mx          int
	ans         int
	start, end  int
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	return node
}

type Pair struct {
	first, second int
}
