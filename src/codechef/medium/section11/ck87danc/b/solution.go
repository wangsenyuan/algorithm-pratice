package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		A := readNNums(reader, n)
		Q := make([][]int, k)
		for i := 0; i < k; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(n, m, k, A, Q)

		for i := 0; i < k; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}

	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const H = 30
const INF = 1 << 30

func solve(n int, m int, k int, A []int, Q [][]int) []int {
	S := make([]int, n+1)
	copy(S[1:], A)
	for i := 1; i <= n; i++ {
		S[i] ^= S[i-1]
	}

	root := NewNode()

	l := make([]int, 0, n)
	r := make([]int, 0, n)

	for i := n; i >= 1; i-- {
		Insert(root, S[i], i)
		j := Search(root, S[i-1], m)
		if j >= INF || len(r) > 0 && r[len(r)-1] <= j {
			continue
		}
		l = append(l, i)
		r = append(r, j)
	}

	reverse(l)
	reverse(r)
	sz := len(l)
	mn := make([][]int, sz+1)

	for i := 0; i <= sz; i++ {
		mn[i] = make([]int, 20)
	}

	for i := 1; i <= sz; i++ {
		mn[i][0] = r[i-1] - l[i-1] + 1
	}

	for j := 1; 1<<uint(j) <= sz; j++ {
		for i := 1; i+(1<<uint(j))-1 <= sz; i++ {
			mn[i][j] = min(mn[i][j-1], mn[i+(1<<uint(j-1))][j-1])
		}
	}

	query := func(l, r int) int {
		d := 31 - clz(r-l+1)
		return min(mn[l][d], mn[r-(1<<uint(d))+1][d])
	}

	ans := make([]int, k)

	for i := 0; i < k; i++ {
		x, y := Q[i][0], Q[i][1]
		st := sort.Search(sz, func(j int) bool {
			return l[j] >= x
		})
		st++
		en := sort.Search(sz, func(j int) bool {
			return r[j] >= y+1
		})

		if st > en {
			ans[i] = -1
		} else {
			ans[i] = query(st, en)
		}
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}

func clz(num int) int {
	var cnt int
	for i := 31; i >= 0; i-- {
		if (num>>uint(i))&1 > 0 {
			return cnt
		}
		cnt++
	}
	return H
}

type Node struct {
	children [2]*Node
	val      int
}

func NewNode() *Node {
	node := new(Node)
	node.val = INF
	return node
}

func Insert(root *Node, x int, id int) {
	node := root

	for i := H - 1; i >= 0; i-- {
		bit := (x >> uint(i)) & 1
		if node.children[bit] == nil {
			node.children[bit] = NewNode()
		}
		node = node.children[bit]
		node.val = min(node.val, id)
	}
}

func Search(root *Node, x, y int) int {
	node := root

	ans := INF

	for i := H - 1; i >= 0; i-- {
		bit := ((x ^ y) >> uint(i)) & 1
		if (y>>uint(i))&1 == 0 {
			if node.children[bit^1] != nil {
				ans = min(ans, node.children[bit^1].val)
			}
		}
		node = node.children[bit]
		if node == nil {
			break
		}
	}

	if node != nil {
		ans = min(ans, node.val)
	}

	return ans
}
