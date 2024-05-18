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

		a := readNNums(reader, n)

		res := solve1(a)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve1(a []int) int {
	n := len(a)
	if n <= 2 {
		return 0
	}

	x, y := n+1, n+1

	var res int
	for i := 0; i < n; i++ {
		if a[i] <= x {
			x = a[i]
		} else if y < a[i] {
			res++
			x = a[i]
		} else {
			y = a[i]
		}
		x, y = min(x, y), max(x, y)
	}

	return res
}

func solve(a []int) int {
	tr := NewNode()

	n := len(a)

	for i := 1; i <= n; i++ {
		ndp := min(tr.Get(1, a[i-1]-1, 1, n)+1, tr.Get(a[i-1], n, 1, n))
		if i > 1 {
			if a[i-2] < a[i-1] {
				tr.Update(1, n, 1, n, 1)
			}
			dp := tr.Get(a[i-2], a[i-2], 1, n)
			if ndp < dp {
				tr.Update(a[i-2], a[i-2], 1, n, ndp-dp)
			}
		}
	}

	return tr.Get(1, n, 1, n)
}

type Node struct {
	left, right *Node
	val         int
	lazy        int
}

func NewNode() *Node {
	node := new(Node)
	return node
}

func (node *Node) pushValue(v int) {
	node.val += v
	node.lazy += v
}

func (node *Node) push(l int, r int) {
	if l == r {
		return
	}
	if node.left == nil {
		node.left = NewNode()
		node.right = NewNode()
	}
	if node.lazy != 0 {
		node.left.pushValue(node.lazy)
		node.right.pushValue(node.lazy)
		node.lazy = 0
	}
}

func (node *Node) pull() {
	node.val = min(node.left.val, node.right.val)
}

func (node *Node) Update(L int, R int, l int, r int, v int) {
	if R < l || r < L {
		return
	}
	node.push(l, r)
	if L <= l && r <= R {
		node.pushValue(v)
		return
	}
	mid := (l + r) / 2
	node.left.Update(L, R, l, mid, v)
	node.right.Update(L, R, mid+1, r, v)
	node.pull()
}

func (node *Node) Get(L int, R int, l int, r int) int {
	if node == nil || R < l || r < L {
		return 1 << 30
	}
	node.push(l, r)

	if L <= l && r <= R {
		return node.val
	}

	mid := (l + r) / 2
	a := node.left.Get(L, R, l, mid)
	b := node.right.Get(L, R, mid+1, r)
	return min(a, b)
}
