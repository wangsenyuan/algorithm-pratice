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
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

func solve(n int, A []int) int {
	var res int

	root := NewNode()

	for i := 1; i < n; i++ {
		if i > 1 {
			res = max(res, FindXorMax(root, A[i]))
		}
		for j := 0; j < i; j++ {
			root.Add(A[j]^A[i], MAX_POS)
		}
	}

	return res
}

const MAX_POS = 29

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	children []*Node
}

func NewNode() *Node {
	node := new(Node)
	node.children = make([]*Node, 2)
	return node
}

func (this *Node) Add(num int, pos uint) {
	x := (num >> pos) & 1
	if this.children[x] == nil {
		this.children[x] = NewNode()
	}
	if pos > 0 {
		this.children[x].Add(num, pos-1)
	}
}

func FindXorMax(root *Node, num int) int {
	cur := root
	var res int
	var pos uint = MAX_POS
	for cur != nil {
		x := (num >> pos) & 1
		if cur.children[1-x] != nil {
			res |= 1 << pos
			cur = cur.children[1-x]
		} else {
			cur = cur.children[x]
		}

		pos--
	}

	return res
}
