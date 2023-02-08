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
		n, m := readTwoNums(reader)
		board := make([][]int, n)
		for i := 0; i < n; i++ {
			board[i] = readNNums(reader, m)
		}
		res := solve(board)
		buf.WriteString(fmt.Sprintf("%d\n", res))
		if tc > 0 {
			readString(reader)
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

const X = 1e6 + 10

func solve(board [][]int) int {
	n := len(board)
	m := len(board[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	jump := func(i, j int) {
		dp[i][j] = 1
		var x, y int = X, 0

		for c := j + 1; c < m; c++ {
			if board[i][j] <= board[i][c] && board[i][c] < x {
				x = board[i][c]
				if dp[i][j] < 1+dp[i][c] {
					dp[i][j] = 1 + dp[i][c]
				}
				if x == board[i][j] {
					break
				}
			}
		}

		for r := i + 1; r < n; r++ {
			if board[i][j] >= board[r][j] && board[r][j] > y {
				y = board[r][j]
				if dp[i][j] < 1+dp[r][j] {
					dp[i][j] = 1 + dp[r][j]
				}
				if board[i][j] == y {
					break
				}
			}
		}
	}
	var res int
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			jump(i, j)
			res = max(res, dp[i][j])
		}
	}
	return res
}

func solve1(board [][]int) int {
	//dp[i][j] = 到达(i,j)的最大值
	//dp[i][j] = some r < i and board[r][j] <= board[i][j] some c < j and board[i][c] <= board[i][j]
	//迭代r & c 会增加O(n)的时间, 如何快速找到 <= board[i][j]中的最大值呢？
	// segment tree?
	n := len(board)
	m := len(board[0])

	cols := make([]*Node, m)
	for j := 0; j < m; j++ {
		cols[j] = NewNode(0, X)
	}
	var best int

	for i := 0; i < n; i++ {
		row := NewNode(0, X)
		for j := 0; j < m; j++ {
			var v int
			if j > 0 {
				v = row.Query(0, board[i][j])
			}
			if i > 0 {
				v = max(v, cols[j].Query(board[i][j], X))
			}
			v++
			best = max(best, v)
			row.Set(board[i][j], v)
			cols[j].Set(board[i][j], v)
		}
	}

	return best
}

type Node struct {
	left  *Node
	right *Node
	start int
	end   int
	val   int
}

func NewNode(start, end int) *Node {
	node := new(Node)
	node.start = start
	node.end = end
	node.val = 0
	return node
}

func (node *Node) GetValue() int {
	if node == nil {
		return 0
	}
	return node.val
}

func (node *Node) push() {
	if node.left != nil {
		return
	}
	mid := (node.start + node.end) / 2
	node.left = NewNode(node.start, mid)
	node.right = NewNode(mid+1, node.end)
}

func (node *Node) Set(p int, v int) {
	if node.start == node.end {
		// == p
		node.val = max(node.val, v)
		return
	}
	node.push()
	mid := (node.start + node.end) / 2
	if p <= mid {
		node.left.Set(p, v)
	} else {
		node.right.Set(p, v)
	}
	node.val = max(node.left.GetValue(), node.right.GetValue())
}

func (node *Node) Query(l int, r int) int {
	if node == nil || r < node.start || node.end < l {
		return 0
	}
	if l <= node.start && node.end <= r {
		return node.val
	}
	a := node.left.Query(l, r)
	b := node.right.Query(l, r)
	return max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
