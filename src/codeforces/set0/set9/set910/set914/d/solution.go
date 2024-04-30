package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	cmds := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var t int
		pos := readInt(s, 0, &t) + 1
		if t == 1 {
			cmds[i] = make([]int, 4)
		} else {
			cmds[i] = make([]int, 3)
		}
		cmds[i][0] = t
		for j := 1; j < len(cmds[i]); j++ {
			pos = readInt(s, pos, &cmds[i][j]) + 1
		}
	}

	res := solve(a, cmds)

	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(a []int, cmds [][]int) []bool {
	tr := Build(a)

	n := len(a)

	var ans []bool

	for _, cmd := range cmds {
		if cmd[0] == 1 {
			l, r, x := cmd[1], cmd[2], cmd[3]
			cnt := tr.Query(l-1, r-1, x, n)
			ans = append(ans, cnt >= r-l)
		} else {
			i, y := cmd[1], cmd[2]
			tr.Update(i-1, y, n)
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Node struct {
	left, right *Node
	val         int
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
			node.val = gcd(node.left.val, node.right.val)
		}
		return node
	}
	return loop(0, len(a)-1)
}

func (tr *Node) Query(L int, R int, x int, n int) int {
	var loop func(node *Node, l int, r int) int

	loop = func(node *Node, l int, r int) int {
		if R < l || r < L {
			return 0
		}

		mid := (l + r) / 2

		if L <= l && r <= R {
			if node.val%x == 0 {
				return r - l + 1
			}
			// node.val % x != 0
			if l == r || node.left.val%x != 0 && node.right.val%x != 0 {
				return 0
			}
			// 至少有一半是ok的
			if node.left.val%x == 0 {
				return mid - l + 1 + loop(node.right, mid+1, r)
			}
			return r - mid + loop(node.left, l, mid)
		}
		return loop(node.left, l, mid) + loop(node.right, mid+1, r)
	}

	return loop(tr, 0, n-1)
}

func (tr *Node) Update(p int, v int, n int) {
	var loop func(node *Node, l int, r int)
	loop = func(node *Node, l int, r int) {
		if l == r {
			node.val = v
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(node.left, l, mid)
		} else {
			loop(node.right, mid+1, r)
		}
		node.val = gcd(node.left.val, node.right.val)
	}

	loop(tr, 0, n-1)
}
