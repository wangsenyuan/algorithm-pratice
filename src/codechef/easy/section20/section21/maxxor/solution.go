package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const H = 31

func solve(A []int) int {

	var D int

	for _, num := range A {
		D ^= num
	}

	n := len(A)

	var xor int
	var best int

	root := NewNode(H - 1)

	root.Add(D)

	for i := 0; i < n; i++ {
		xor ^= A[i]
		cur := xor | D

		node := root

		var tmp int

		for node != nil {
			x := (cur >> uint(node.d)) & 1
			// to maximize the result, better to go next path
			if node.children[1^x] != nil {
				tmp |= 1 << uint(node.d)
				node = node.children[1^x]
			} else {
				node = node.children[x]
			}
		}

		root.Add(cur)

		best = max(best, tmp)
	}

	return best ^ D
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	children []*Node
	d        int
}

func NewNode(d int) *Node {
	node := new(Node)
	node.children = make([]*Node, 2)
	node.d = d
	return node
}

func (node *Node) Add(num int) {
	if node.d < 0 {
		return
	}
	x := (num >> uint(node.d)) & 1
	if node.children[x] == nil {
		node.children[x] = NewNode(node.d - 1)
	}
	node.children[x].Add(num)
}
