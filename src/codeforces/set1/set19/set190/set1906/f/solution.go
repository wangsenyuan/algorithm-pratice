package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	operations := make([][]int, m)
	for i := 0; i < m; i++ {
		operations[i] = readNNums(reader, 3)
	}
	k := readNum(reader)
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		queries[i] = readNNums(reader, 3)
	}
	res := solve(n, operations, queries)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(n int, operations [][]int, queries [][]int) []int {
	m := len(operations)
	qn := len(queries)
	qr := make([][]int, n)

	for i, cur := range queries {
		k := cur[0] - 1
		qr[k] = append(qr[k], i)
	}

	open := make([][]int, n)
	end := make([][]int, n)

	for i, op := range operations {
		l, r := op[0]-1, op[1]-1
		open[l] = append(open[l], i)
		end[r] = append(end[r], i)
	}

	tr := NewTree(m)

	ans := make([]int, qn)

	for k := 0; k < n; k++ {
		for _, i := range open[k] {
			x := operations[i][2]
			tr.Update(i, x, true)
		}

		for _, i := range qr[k] {
			l, r := queries[i][1]-1, queries[i][2]
			tmp := tr.Get(l, r)
			if tmp.cnt < r-l {
				// 如果k没有被l...r的所有的操作覆盖，那么它可以取值为0
				// 否则的话，它必须根据这些操作的最小值（有可能是负数)
				tmp.f00 = max(tmp.f00, 0)
			}
			ans[i] = tmp.f00
		}

		for _, i := range end[k] {
			x := operations[i][2]
			tr.Update(i, x, false)
		}
	}

	return ans
}

type Node struct {
	f00 int
	f01 int
	f10 int
	f11 int
	cnt int
}

func (node *Node) updateValue(v, w int) {
	node.f00 = v
	node.f01 = v
	node.f10 = v
	node.f11 = v
	node.cnt = w
}

func NewNode(v int) *Node {
	node := new(Node)
	node.f00 = v
	node.f01 = v
	node.f10 = v
	node.f11 = v
	return node
}

func unit() *Node {
	node := NewNode(-inf)
	node.f11 = 0
	node.cnt = 0
	return node
}

const inf = 1 << 50

func maintain(left, right *Node) *Node {
	node := unit()
	node.f00 = max(-inf, left.f00, right.f00, left.f01+right.f10)
	node.f10 = max(-inf, left.f10, left.f11+right.f10)
	node.f01 = max(-inf, right.f01, left.f01+right.f11)
	node.f11 = left.f11 + right.f11
	node.cnt = left.cnt + right.cnt
	return node
}

type Tree struct {
	nodes []*Node
	sz    int
}

func NewTree(n int) *Tree {
	arr := make([]*Node, 2*n)
	for i := 0; i < 2*n; i++ {
		arr[i] = unit()
	}

	return &Tree{arr, n}
}

func (tr *Tree) Update(pos int, v int, add bool) {
	pos += tr.sz

	if add {
		tr.nodes[pos].updateValue(v, 1)
	} else {
		tr.nodes[pos].updateValue(-inf, 0)
		tr.nodes[pos].f11 = 0
	}

	for pos > 1 {
		pos >>= 1
		tr.nodes[pos] = maintain(tr.nodes[pos*2], tr.nodes[pos*2+1])
	}
}

func (tr *Tree) Get(l int, r int) *Node {
	l += tr.sz
	r += tr.sz
	// 这边左右是有关系的
	ra := unit()
	rb := unit()

	for l < r {
		if l&1 == 1 {
			ra = maintain(ra, tr.nodes[l])
			l++
		}
		if r&1 == 1 {
			r--
			rb = maintain(tr.nodes[r], rb)
		}
		l >>= 1
		r >>= 1
	}

	return maintain(ra, rb)
}
