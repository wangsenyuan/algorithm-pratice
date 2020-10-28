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

	n, q := readTwoNums(reader)
	P := readNNums(reader, n)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(reader, 2)
	}
	res := solve(n, P, q, queries)
	fmt.Print(res)
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

func solve(n int, P []int, q int, queries [][]int) string {
	sort.Ints(P)
	var node *Node

	for i := 0; i < n; i++ {
		node = node.Set(P[i])
	}

	get := func() int {
		if node == nil {
			// empty node
			return 0
		}
		res := node.maxVal - node.minVal - node.maxGap
		return int(res)
	}

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", get()))

	for i := 0; i < q; i++ {
		cur := queries[i]
		if cur[0] == 0 {
			node = node.Unset(cur[1])
		} else {
			node = node.Set(cur[1])
		}
		buf.WriteString(fmt.Sprintf("%d\n", get()))

	}

	return buf.String()
}

const MAX_X = 1000000010
const INF = 1 << 28
const N_INF = -INF

type Node struct {
	left, right *Node
	start, end  int
	maxGap      int64
	minVal      int64
	maxVal      int64
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.maxGap = -1
	node.minVal = INF
	node.maxVal = N_INF

	return node
}

func (node *Node) merge() *Node {
	if node.left == nil && node.right == nil {
		return nil
	}

	if node.right == nil {
		node.minVal = node.left.minVal
		node.maxVal = node.left.maxVal
		node.maxGap = node.left.maxGap
	} else if node.left == nil {
		node.minVal = node.right.minVal
		node.maxVal = node.right.maxVal
		node.maxGap = node.right.maxGap
	} else {
		node.minVal = min(node.left.minVal, node.right.minVal)
		node.maxVal = max(node.left.maxVal, node.right.maxVal)
		node.maxGap = max(node.right.minVal-node.left.maxVal, max(node.left.maxGap, node.right.maxGap))
	}

	return node
}

func (node *Node) Unset(p int) *Node {
	if node.start+1 == node.end {
		// node.start == p
		return nil
	}

	mid := (node.start + node.end) / 2

	if p < mid {
		node.left = node.left.Unset(p)
	} else {
		node.right = node.right.Unset(p)
	}

	return node.merge()
}

func (root *Node) Set(p int) *Node {
	var loop func(node *Node, start, end int) *Node

	loop = func(node *Node, start, end int) *Node {
		if node == nil {
			node = NewNode(start, end)
		}
		if start+1 == end {
			node.minVal = int64(p)
			node.maxVal = int64(p)
			node.maxGap = 0
			return node
		}

		mid := (start + end) / 2

		if p < mid {
			node.left = loop(node.left, start, mid)
		} else {
			node.right = loop(node.right, mid, end)
		}

		return node.merge()
	}

	return loop(root, 0, MAX_X)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
