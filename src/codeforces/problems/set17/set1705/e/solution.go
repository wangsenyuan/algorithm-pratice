package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	n, m := readTwoNums(reader)
	A := readNNums(reader, n)
	ques := make([][]int, m)
	for i := 0; i < m; i++ {
		ques[i] = readNNums(reader, 2)
	}
	res := solve(A, ques)

	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
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

const MAX_A = 200100

type Node struct {
	left, right *Node
	start, end  int
	tag         int
	val         int
	lazy        bool
}

func NewNode(start int, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.tag = 0
	node.val = 0
	node.lazy = false
	if start < end {
		mid := (start + end) / 2
		node.left = NewNode(start, mid)
		node.right = NewNode(mid+1, end)
	}
	return node
}

func (node *Node) push() {
	if !node.lazy {
		return
	}

	node.val = (node.end - node.start + 1) * node.tag

	if node.start == node.end {
		return
	}

	node.left.lazy = node.lazy
	node.left.tag = node.tag
	node.right.lazy = node.lazy
	node.right.tag = node.tag
	node.tag = 0
	node.lazy = false
}

func (node *Node) Update(start int, end int, x int) {
	node.push()
	if node.end < start || end < node.start {
		// out of bound
		return
	}
	if start <= node.start && node.end <= end {
		node.tag = x
		node.lazy = true
		node.push()
	} else {
		node.left.Update(start, end, x)
		node.right.Update(start, end, x)
		node.val = node.left.val + node.right.val
	}
}

func (node *Node) Query(start int, end int) int {
	if node.end < start || end < node.start {
		return 0
	}
	node.push()

	if start <= node.start && node.end <= end {
		return node.val
	}

	if node.left == nil {
		return 0
	}

	return node.left.Query(start, end) + node.right.Query(start, end)
}

// pre = prefix_sum([0...k))
func (node *Node) max_right(k int, pre int, v int) int {
	node.push()

	if node.start == node.end {
		if node.val == v {
			return node.start
		}
		return node.start - 1
	}

	node.left.push()

	mid := (node.start + node.end) / 2
	if mid < k {
		return node.right.max_right(k, pre-node.left.val, v)
	}
	if node.left.val-pre == v*(mid+1-k) {
		return node.right.max_right(mid+1, 0, v)
	}

	return node.left.max_right(k, pre, v)
}

func (node *Node) get_answer() int {
	node.push()
	if node.start == node.end {
		if node.val == 1 {
			return node.start
		}
		return node.start - 1
	}

	node.right.push()

	if node.right.val == 0 {
		return node.left.get_answer()
	}

	return node.right.get_answer()
}

func solve(A []int, queries [][]int) []int {
	root := NewNode(0, MAX_A)

	add := func(x int) {
		// x 右边连续1的位置
		y := root.max_right(x, root.Query(0, x-1), 1) + 1
		if x == y {
			root.Update(x, x, 1)
		} else {
			root.Update(x, y-1, 0)
			root.Update(y, y, 1)
		}
	}

	remove := func(x int) {
		y := root.max_right(x, root.Query(0, x-1), 0) + 1
		if x == y {
			root.Update(x, x, 0)
		} else {
			root.Update(x, y-1, 1)
			root.Update(y, y, 0)
		}
	}

	for i := 0; i < len(A); i++ {
		add(A[i])
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		k, l := cur[0], cur[1]
		k--
		remove(A[k])
		add(l)
		A[k] = l
		ans[i] = root.get_answer()
	}
	return ans
}
