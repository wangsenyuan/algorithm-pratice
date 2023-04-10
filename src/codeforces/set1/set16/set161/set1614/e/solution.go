package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//var buf bytes.Buffer
	n := readNum(reader)
	T := make([]int, n)
	X := make([][]int, n)
	for i := 0; i < n; i++ {
		T[i] = readNum(reader)
		k := readNum(reader)
		X[i] = readNNums(reader, k)
	}
	res := solve(n, T, X)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, T []int, X [][]int) []int {
	root := NewNode()

	var v int
	var ans int

	var res []int

	for i := 0; i < n; i++ {
		ti := T[i]
		if ti >= v {
			erase(root, 0, INF, ti-v)
		} else {
			v--
		}

		if ti > v {
			erase(root, 0, INF, ti-1-v)
			v++
		}
		for _, x := range X[i] {
			x = (x + ans) % 1000000001
			ans = v + x - getSum(root, 0, 1<<30, x)
			res = append(res, ans)
		}
	}

	return res
}

var INF int = 1 << 30

type Node struct {
	left, right *Node
	sum         int
}

func NewNode() *Node {
	n := new(Node)
	n.sum = 0
	return n
}

func (node *Node) getSum() int {
	if node == nil {
		return 0
	}
	return node.sum
}

func erase(node *Node, l int, r int, k int) *Node {
	if node == nil {
		node = NewNode()
	}
	node.sum++

	if r-l == 1 {
		return node
	}

	mid := (l + r) / 2

	tmp := mid - l - node.left.getSum()

	if k < tmp {
		node.left = erase(node.left, l, mid, k)
	} else {
		node.right = erase(node.right, mid, r, k-tmp)
	}

	return node
}

func getSum(node *Node, l int, r int, x int) int {
	if x <= l || node == nil {
		return 0
	}
	if x >= r {
		return node.sum
	}
	mid := (l + r) / 2
	return getSum(node.left, l, mid, x) + getSum(node.right, mid, r, x)
}
