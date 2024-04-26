package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	k, m := readTwoNums(reader)
	a := readNNums(reader, 1<<k)
	cmds := make([][]int, m)
	for i := 0; i < m; i++ {
		var t int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &t) + 1
		if t == 1 || t == 4 {
			cmds[i] = make([]int, 3)
		} else {
			cmds[i] = make([]int, 2)
		}
		cmds[i][0] = t
		for j := 1; j < len(cmds[i]); j++ {
			pos = readInt(s, pos, &cmds[i][j]) + 1
		}
	}

	res := solve(k, a, cmds)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Println(buf.String())
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

func solve(k int, a []int, cmds [][]int) []int {
	n := len(a)
	tr := Build(a)

	swap := make([]bool, k+1)

	var res []int
	for _, cmd := range cmds {
		if cmd[0] == 1 {
			i, x := cmd[1]-1, cmd[2]
			tr.update(i, x, n, swap, k)
		} else if cmd[0] == 2 {
			for i := 0; i <= cmd[1]; i++ {
				swap[i] = !swap[i]
			}
		} else if cmd[0] == 3 {
			i := cmd[1]
			swap[i+1] = !swap[i+1]
		} else {
			l, r := cmd[1], cmd[2]

			res = append(res, tr.query(l-1, r-1, n, swap, k))
		}
	}

	return res
}

type Node struct {
	left  *Node
	right *Node
	val   int
}

func Build(a []int) *Node {
	var loop func(l int, r int) *Node
	loop = func(l int, r int) *Node {
		node := new(Node)
		if l == r {
			node.val = a[l]
		} else {
			mid := (l + r) / 2
			node.left = loop(l, mid)
			node.right = loop(mid+1, r)
			node.val = node.left.val + node.right.val
		}
		return node
	}

	return loop(0, len(a)-1)
}

func (root *Node) update(p int, v int, n int, swap []bool, d int) {
	var loop func(node *Node, l int, r int, d int)
	loop = func(node *Node, l int, r int, d int) {
		if l == r {
			node.val = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			next := node.left
			if swap[d] {
				next = node.right
			}
			loop(next, l, mid, d-1)
		} else {
			next := node.right
			if swap[d] {
				next = node.left
			}
			loop(next, mid+1, r, d-1)
		}
		node.val = node.left.val + node.right.val
	}

	loop(root, 0, n-1, d)
}

func (root *Node) query(L int, R int, n int, swap []bool, d int) int {
	var loop func(node *Node, l int, r int, d int) int
	loop = func(node *Node, l int, r int, d int) int {
		if R < l || r < L {
			return 0
		}
		if L <= l && r <= R {
			return node.val
		}
		mid := (l + r) / 2
		var res int
		if L <= mid {
			next := node.left
			if swap[d] {
				next = node.right
			}
			res += loop(next, l, mid, d-1)
		}
		if mid < R {
			next := node.right
			if swap[d] {
				next = node.left
			}
			res += loop(next, mid+1, r, d-1)
		}
		return res
	}

	return loop(root, 0, n-1, d)
}
