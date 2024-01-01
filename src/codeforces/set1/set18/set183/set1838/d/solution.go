package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	s := readString(reader)[:n]
	queries := make([]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNum(reader)
	}
	res := solve(s, queries)
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

func solve(s string, queries []int) []bool {
	// 假设从0可以按照正常进行到i,i+1 = ')',显然无法继续
	// 如果 在i前面有连续两个 '(', 那么就可以到达下一个位置
	// ()()()()), 也就是说只要有两个连续的 (, 就肯定能到达最后的位置
	// 如果不存在，那么就必须是()的重复
	// 0 - indexed the set contains odd i, but s[i] = '('
	// and even i, but s[i] = ')'
	var root *Node

	insert := func(x int) {
		var L, R *Node
		Split(root, x, &L, &R)
		cur := NewNode(x)
		root = Merge(Merge(L, cur), R)
	}

	remove := func(x int) {
		var L, R, p *Node
		Split(root, x, &L, &R)
		Split(L, x-1, &L, &p)
		// remove x
		if p != nil {
			p = Merge(p.ls, p.rs)
		}
		root = Merge(Merge(L, p), R)
	}

	buf := []byte(s)

	for i := 0; i < len(s); i++ {
		if i&1 == 0 && buf[i] == ')' {
			insert(i)
		} else if i&1 == 1 && buf[i] == '(' {
			insert(i)
		}
	}

	n := len(s)

	flip := func(x int) bool {
		if n&1 == 1 {
			return false
		}
		if x&1 == 0 && buf[x] == ')' || x&1 == 1 && buf[x] == '(' {
			remove(x)
		}
		// do flip
		if buf[x] == '(' {
			buf[x] = ')'
		} else {
			buf[x] = '('
		}

		if x&1 == 0 && buf[x] == ')' || x&1 == 1 && buf[x] == '(' {
			insert(x)
		}

		if root == nil {
			return true
		}
		a := MinValue(root)
		b := MaxValue(root)
		if a&1 == 0 || b&1 == 1 {
			return false
		}
		return true
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		ans[i] = flip(cur - 1)
	}

	return ans
}

type Node struct {
	ls, rs *Node
	key    int
	pri    int
	size   int
	lazy   int
}

func NewNode(x int) *Node {
	n := new(Node)
	n.key = x
	n.size = 1
	n.pri = rand.Int()
	return n
}

func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	return node.size
}

func (node *Node) update() {
	node.size = node.ls.Size() + node.rs.Size() + 1
}

func Split(root *Node, x int, L **Node, R **Node) {
	if root == nil {
		*L = nil
		*R = nil
		return
	}
	if root.key <= x {
		*L = root
		Split(root.rs, x, &(root.rs), R)
	} else {
		*R = root
		Split(root.ls, x, L, &(root.ls))
	}
	root.update()
}

func Merge(L *Node, R *Node) *Node {
	if L == nil || R == nil {
		if R != nil {
			return R
		}
		return L
	}

	if L.pri > R.pri {
		L.rs = Merge(L.rs, R)
		L.update()
		return L
	}
	R.ls = Merge(L, R.ls)
	R.update()
	return R
}

func MinValue(root *Node) int {
	if root.ls != nil {
		return MinValue(root.ls)
	}
	return root.key
}

func MaxValue(root *Node) int {
	if root.rs != nil {
		return MaxValue(root.rs)
	}
	return root.key
}
