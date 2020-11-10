package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, u := readTwoNums(reader)

	inst := make([][]int, u)
	for i := 0; i < u; i++ {
		inst[i] = readNNums(reader, 2)
	}

	fmt.Println(solve(n, u, inst))
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

func init() {
	rand.Seed(41)
}

func solve(n, u int, inst [][]int) string {
	arr := make([]uint64, n)

	for i := 0; i < n; i++ {
		arr[i] = uint64(rand.Int63())
		for arr[i] == 0 {
			arr[i] = uint64(rand.Int63())
		}
	}

	nodes := make([]*Node, u+1)
	nodes[0] = Build(arr, 0, n)
	var cur int
	for i := 1; i <= u; i++ {
		l, r := inst[i-1][0], inst[i-1][1]
		l--
		r--
		nodes[i] = nodes[i-1].Update(l, r+1, 0, n)
		if small(nodes[cur], nodes[i], 0, n) {
			cur = i
		}
	}

	var buf bytes.Buffer

	var dfs func(node *Node, l, r int)
	dfs = func(node *Node, l, r int) {
		if l+1 == r {
			if node.h1 > 0 {
				buf.WriteByte('1')
			} else {
				buf.WriteByte('0')
			}
			return
		}

		mid := (l + r) / 2
		node = node.Push()
		dfs(node.left, l, mid)
		dfs(node.right, mid, r)
	}

	dfs(nodes[cur], 0, n)

	return buf.String()
}

type Node struct {
	left, right *Node
	h0, h1      uint64
	swap        int
}

func NewNode(l, r *Node, h0, h1 uint64, swap int) *Node {
	node := new(Node)
	node.left = l
	node.right = r
	node.h0 = h0
	node.h1 = h1
	node.swap = swap
	return node
}

func (node *Node) Push() *Node {
	if node.swap == 0 {
		return node
	}
	var x, y *Node

	if node.left != nil {
		x = NewNode(node.left.left, node.left.right, node.left.h1, node.left.h0, node.left.swap^1)
	}

	if node.right != nil {
		y = NewNode(node.right.left, node.right.right, node.right.h1, node.right.h0, node.right.swap^1)
	}

	return NewNode(x, y, node.h0, node.h1, 0)
}

func (node *Node) Update(l, r int, tl, tr int) *Node {
	if tl >= r || tr <= l {
		return node
	}
	if tl >= l && tr <= r {
		return NewNode(node.left, node.right, node.h1, node.h0, node.swap^1)
	}

	mid := (tl + tr) / 2
	node = node.Push()
	a := node.left.Update(l, r, tl, mid)
	b := node.right.Update(l, r, mid, tr)
	return NewNode(a, b, a.h0^b.h0, a.h1^b.h1, node.swap)
}

func small(a, b *Node, tl, tr int) bool {
	if tr-tl == 1 {
		return a.h0 > 0 && b.h1 > 0
	}

	a = a.Push()
	b = b.Push()

	mid := (tl + tr) / 2
	if a.left.h0 == b.left.h0 {
		return small(a.right, b.right, mid, tr)
	}
	return small(a.left, b.left, tl, mid)
}

func Build(arr []uint64, l, r int) *Node {
	if l+1 == r {
		return NewNode(nil, nil, arr[l], 0, 0)
	}
	mid := (l + r) / 2
	a := Build(arr, l, mid)
	b := Build(arr, mid, r)
	return NewNode(a, b, a.h0^b.h0, a.h1^b.h1, 0)
}
