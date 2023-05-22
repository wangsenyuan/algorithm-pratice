package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		s := readString(reader)[:n]
		pos := make([]int, q)
		x := make([]byte, q)
		for i := 0; i < q; i++ {
			line, _ := reader.ReadBytes('\n')
			j := readInt(line, 0, &pos[i])
			x[i] = line[j+1]
		}
		res := solve(s, pos, x)
		for i := 0; i < q; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

func solve(s string, pos []int, C []byte) []int {
	tree := BuildTree(s)

	n := len(pos)

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		j := pos[i] - 1
		x := C[i]
		tree.Update(j, x)
		ans[i] = tree.nodes[1].val[5]
	}

	return ans
}

// a, b, c, ab, bc, abc
type Node struct {
	val [6]int
}

func NewNode(x byte) *Node {
	node := new(Node)
	node.val[int(x-'a')]++
	return node
}

func (this *Node) Merge(that *Node) *Node {
	res := new(Node)
	// 在节点中包含的a的个数（要删除的个数)
	for i := 0; i < 3; i++ {
		res.val[i] = this.val[i] + that.val[i]
	}
	// ab的个数  (a + ab)  (ab + b)
	res.val[3] = min(this.val[0]+that.val[3], this.val[3]+that.val[1])
	// bc的个数  (b + bc)  (bc + c)
	res.val[4] = min(this.val[1]+that.val[4], this.val[4]+that.val[2])
	// a + abc, ab + bc, abc + c
	res.val[5] = min(this.val[0]+that.val[5], min(this.val[3]+that.val[4], this.val[5]+that.val[2]))

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

type Tree struct {
	nodes []*Node
}

func BuildTree(s string) *Tree {
	n := len(s)
	nodes := make([]*Node, 4*n)
	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if l+1 == r {
			nodes[i] = NewNode(s[l])
			return
		}
		mid := (l + r) / 2
		loop(2*i, l, mid)
		loop(2*i+1, mid, r)
		nodes[i] = nodes[2*i].Merge(nodes[2*i+1])
	}

	loop(1, 0, n)

	return &Tree{nodes}
}

func (t *Tree) Update(p int, x byte) {
	n := len(t.nodes) / 4

	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if l+1 == r {
			t.nodes[i] = NewNode(x)
			return
		}
		mid := (l + r) / 2

		if p < mid {
			loop(2*i, l, mid)
		} else {
			loop(2*i+1, mid, r)
		}
		t.nodes[i] = t.nodes[2*i].Merge(t.nodes[2*i+1])
	}

	loop(1, 0, n)
}
