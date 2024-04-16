package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k, m := readThreeNums(reader)
	tariffs := make([][]int, m)
	for i := 0; i < m; i++ {
		tariffs[i] = readNNums(reader, 4)
	}

	res := solve(n, tariffs, k)

	fmt.Println(res)
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

func solve(n int, tariffs [][]int, k int) int {
	var p int
	for _, tariff := range tariffs {
		p = max(p, tariff[3])
	}
	p++

	activate := make([][]int, n+3)
	deactivate := make([][]int, n+3)
	for i := 0; i <= n+2; i++ {
		activate[i] = make([]int, 0, 1)
		deactivate[i] = make([]int, 0, 1)
	}

	for i, tariff := range tariffs {
		l, r := tariff[0], tariff[1]
		activate[l] = append(activate[l], i)
		deactivate[r+1] = append(deactivate[r+1], i)
	}

	tr := Build(p)

	var res int

	for i := 1; i <= n; i++ {
		for _, j := range deactivate[i] {
			tariff := tariffs[j]
			c, x := tariff[2], tariff[3]
			tr.Update(x, -c, p)
		}
		for _, j := range activate[i] {
			tariff := tariffs[j]
			c, x := tariff[2], tariff[3]
			tr.Update(x, c, p)
		}
		res += tr.Query(k, p)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	left, right *Node
	cnt         int
	val         int
}

func Build(n int) *Node {
	var loop func(l int, r int) *Node
	loop = func(l int, r int) *Node {
		node := new(Node)
		node.cnt = 0
		if l == r {
			return node
		}
		mid := (l + r) / 2
		node.left = loop(l, mid)
		node.right = loop(mid+1, r)
		return node
	}

	return loop(0, n-1)
}

func (root *Node) Update(p int, v int, n int) {
	var loop func(node *Node, l int, r int)
	loop = func(node *Node, l int, r int) {
		if l == r {
			node.cnt += v
			node.val = l * node.cnt
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(node.left, l, mid)
		} else {
			loop(node.right, mid+1, r)
		}
		node.cnt = node.left.cnt + node.right.cnt
		node.val = node.left.val + node.right.val
	}

	loop(root, 0, n-1)
}

func (root *Node) Query(k int, n int) int {
	var loop func(node *Node, l int, r int, expect int) int
	loop = func(node *Node, l int, r int, expect int) int {
		if node.cnt <= expect {
			return node.val
		}

		if l == r {
			return l * expect
		}

		mid := (l + r) / 2
		// node.cnt > expect
		if node.left.cnt >= expect {
			return loop(node.left, l, mid, expect)
		}

		return node.left.val + loop(node.right, mid+1, r, expect-node.left.cnt)
	}

	return loop(root, 0, n-1, k)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
