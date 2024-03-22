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
	s := readString(reader)[:n]

	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}

	res := solve(s, queries)

	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(s string, queries []string) []int {
	tr := Build(s)

	var ans []int

	for _, cur := range queries {
		if cur == "count" {
			ans = append(ans, tr.value[0])
		} else {
			var l, r int
			pos := readInt([]byte(cur), len("switch")+1, &l)
			readInt([]byte(cur), pos+1, &r)
			l--
			r--
			tr.Flip(l, r, len(s))
		}
	}

	return ans
}

type Node struct {
	flag  bool   // 是否翻转
	value [2]int // 0最长递增序列, 1最长递减序列
	cnt   [2]int // 0/1的数量
	left  *Node
	right *Node
}

func (node *Node) pull() {
	node.value[0] = max(node.left.value[0]+node.right.cnt[1], node.left.cnt[0]+node.right.value[0])
	node.value[1] = max(node.left.value[1]+node.right.cnt[0], node.left.cnt[1]+node.right.value[1])

	node.cnt[0] = node.left.cnt[0] + node.right.cnt[0]
	node.cnt[1] = node.left.cnt[1] + node.right.cnt[1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (node *Node) push() {
	if node.flag {
		if node.left != nil {
			node.left.flip()
			node.right.flip()
		}
		node.flag = false
	}
}

func (node *Node) flip() {
	node.value[0], node.value[1] = node.value[1], node.value[0]
	node.cnt[0], node.cnt[1] = node.cnt[1], node.cnt[0]
	node.flag = !node.flag
}

func Build(s string) *Node {

	var loop func(l int, r int) *Node

	loop = func(l int, r int) *Node {
		node := new(Node)
		if l == r {
			x := checkValue(s[l])
			node.cnt[x]++
			node.value[0] = 1
			node.value[1] = 1
			return node
		}
		mid := (l + r) / 2
		node.left = loop(l, mid)
		node.right = loop(mid+1, r)
		node.pull()
		return node
	}

	return loop(0, len(s)-1)
}

func checkValue(x byte) int {
	if x == '4' {
		return 0
	}
	return 1
}

func (root *Node) Flip(L int, R int, n int) {

	var loop func(node *Node, l int, r int)
	loop = func(node *Node, l int, r int) {
		node.push()

		if r < L || R < l {
			// out of bounds
			return
		}
		if L <= l && r <= R {
			node.flip()
			return
		}
		mid := (l + r) / 2
		loop(node.left, l, mid)
		loop(node.right, mid+1, r)
		node.pull()
	}

	loop(root, 0, n-1)
}
