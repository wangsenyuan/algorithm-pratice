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
	Q := make([]string, m)
	for i := 0; i < m; i++ {
		Q[i] = readString(reader)
	}
	res := solve(n, Q)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}
	fmt.Print(buf.String())
}

func solve(n int, Q []string) []int64 {
	type Node struct {
		left, right *Node
		l, r        int
		color       int
		val         int64
	}

	var build func(l int, r int) *Node

	var updateColor func(node *Node, L int, R int, c int)

	var query func(node *Node, x int, res int64) int64

	build = func(l int, r int) *Node {
		node := new(Node)
		node.l = l
		node.r = r
		node.color = 1
		node.val = 0

		if l == r {
			return node
		}
		mid := (l + r) / 2
		node.left = build(l, mid)
		node.right = build(mid+1, r)
		return node
	}

	sum := make(map[int]int64)

	updateColor = func(node *Node, L int, R int, c int) {
		if R < node.l || node.r < L {
			return
		}
		mid := (node.l + node.r) / 2
		if L == node.l && node.r == R {
			if node.color > 0 {
				node.val += sum[node.color] - sum[c]
				node.color = c
			} else {
				updateColor(node.left, L, mid, c)
				updateColor(node.right, mid+1, R, c)
				node.color = c
			}

			return
		}
		// push color down
		if node.color > 0 {
			node.left.color = node.color
			node.right.color = node.color
			node.color = 0
		}
		if R <= mid {
			updateColor(node.left, L, R, c)
		} else if L > mid {
			updateColor(node.right, L, R, c)
		} else {
			updateColor(node.left, L, mid, c)
			updateColor(node.right, mid+1, R, c)
		}
	}
	var col int

	query = func(node *Node, x int, res int64) int64 {
		if col == 0 && node.color != 0 {
			col = node.color
		}
		res += node.val
		if node.l == node.r {
			return res
		}
		mid := (node.l + node.r) / 2
		if x <= mid {
			return query(node.left, x, res)
		}
		return query(node.right, x, res)
	}

	root := build(0, n-1)

	var res []int64

	for _, cur := range Q {
		buf := []byte(cur)
		if buf[0] == 'C' {
			var l, r, c int
			pos := readInt(buf, len("Color")+1, &l)
			pos = readInt(buf, pos+1, &r)
			readInt(buf, pos+1, &c)
			updateColor(root, l-1, r-1, c)
		} else if buf[0] == 'A' {
			var c, x int
			pos := readInt(buf, len("Add")+1, &c)
			readInt(buf, pos+1, &x)
			sum[c] += int64(x)
		} else {
			var i int
			readInt(buf, len("Query")+1, &i)
			i--
			col = 0
			tmp := query(root, i, 0)
			tmp += sum[col]
			res = append(res, tmp)
		}
	}
	return res
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
