package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readNNums(reader, n)

	m := readNum(reader)
	ops := make([][]int, m)

	for i := 0; i < m; i++ {
		ops[i] = readAtMostNNums(reader, 3)
	}

	res := solve(A, ops)

	var buf bytes.Buffer

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}

	fmt.Print(buf.String())
}

func readAtMostNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, 0, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}

		if x == len(bs) {
			break
		}

		var y int

		x = readInt(bs, x, &y)

		res = append(res, y)
	}

	return res
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

func solve(A []int, ops [][]int) []int64 {

	n := len(A)

	root := NewNode(0, n-1)

	for i := 0; i < n; i++ {
		root.Init(i, i, A[i])
	}

	var res []int64

	for _, op := range ops {
		l, r := op[0], op[1]

		if len(op) == 3 {
			if l <= r {
				root.Update(l, r, op[2])
			} else {
				root.Update(0, r, op[2])
				root.Update(l, n-1, op[2])
			}
		} else {
			if l <= r {
				res = append(res, root.Query(l, r))
			} else {
				// r < l
				a := root.Query(0, r)
				b := root.Query(l, n-1)
				res = append(res, min(a, b))
			}
		}
	}

	return res
}

const INF = int64(1) << 60

type Node struct {
	left, right *Node
	start, end  int
	lazy        int64
	val         int64
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.lazy = 0
	node.val = INF
	return node
}

func (node *Node) push() {
	if node.start == node.end {
		return
	}
	if node.left == nil {
		mid := (node.start + node.end) / 2
		node.left = NewNode(node.start, mid)
		node.right = NewNode(mid+1, node.end)
	}
	if node.lazy != 0 {
		node.left.lazy += node.lazy
		node.left.val += node.lazy
		node.right.lazy += node.lazy
		node.right.val += node.lazy
		node.lazy = 0
	}
}

func (node *Node) Init(L int, R int, V int) {
	if R < node.start || node.end < L {
		// out of range
		return
	}
	node.push()

	if L <= node.start && node.end <= R {
		// totally in range
		node.val = int64(V)
		node.lazy = 0
		return
	}

	node.left.Init(L, R, V)
	node.right.Init(L, R, V)

	node.val = min(node.left.val, node.right.val)
}

func (node *Node) Update(L int, R int, V int) {
	if R < node.start || node.end < L {
		// out of range
		return
	}
	node.push()

	if L <= node.start && node.end <= R {
		// totally in range
		node.val += int64(V)
		node.lazy += int64(V)
		return
	}

	node.left.Update(L, R, V)
	node.right.Update(L, R, V)

	node.val = min(node.left.val, node.right.val)
}

func (node *Node) Query(L int, R int) int64 {
	if R < node.start || node.end < L {
		return INF
	}

	node.push()

	if L <= node.start && node.end <= R {
		return node.val
	}
	a := node.left.Query(L, R)
	b := node.right.Query(L, R)
	return min(a, b)
}

type Num interface {
	int | int64
}

func min[T Num](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
